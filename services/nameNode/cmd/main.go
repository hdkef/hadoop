package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	pkgDragonfly "github.com/hdkef/hadoop/pkg/repository/dragonfly"
	pkgPostgres "github.com/hdkef/hadoop/pkg/repository/postgres"
	pkgTransactionable "github.com/hdkef/hadoop/pkg/repository/transactionable"
	pkgSvc "github.com/hdkef/hadoop/pkg/services/impl"
	"github.com/hdkef/hadoop/services/nameNode/config"
	"github.com/hdkef/hadoop/services/nameNode/delivery/grpc"
	dgImpl "github.com/hdkef/hadoop/services/nameNode/repository/impl/dragonfly"
	pgImpl "github.com/hdkef/hadoop/services/nameNode/repository/impl/postgres"
	svcImpl "github.com/hdkef/hadoop/services/nameNode/service/impl"
	ucImpl "github.com/hdkef/hadoop/services/nameNode/usecase/impl"
	"golang.org/x/sync/errgroup"
)

func main() {
	// init config
	cfg := config.NewConfig()

	// init db
	db := pkgPostgres.NewPostgresConn(cfg.PostgresConfig)
	rdbClient := pkgDragonfly.NewDragonFlyRepo(cfg.DragonFlyConfig)

	// init repo
	iNodeRepo := pgImpl.NewINodeRepo(db)
	transactionRepo := pgImpl.NewTransactionsRepo(db)
	metadataRepo := pgImpl.NewMetadataRepo(db)
	nodeStorageRepo := dgImpl.NewNodeStorage(&rdbClient)

	// init service
	dataNodeCache := make(map[string]*pkgEt.ServiceDiscovery)
	dataNodeCachemtx := &sync.Mutex{}

	dataNodeSvc := svcImpl.NewDataNodeService(cfg)
	nodeAllocatorSvc := svcImpl.NewNodeAllocator(cfg)
	rollbackSvc := svcImpl.NewRollbackService(&svcImpl.RollbackServiceDto{
		DataNodeCache:    dataNodeCache,
		Mtx:              dataNodeCachemtx,
		TransactionsRepo: &transactionRepo,
		DataNodeService:  &dataNodeSvc,
		MetadataRepo:     &metadataRepo,
	})
	serviceRegistry := pkgSvc.NewServiceRegistry(cfg.ServiceRegistryConfig)
	transactionInjector := pkgTransactionable.NewTransactionInjector(db)

	// init usecase

	writeUC := ucImpl.NewWriteUsecase(&ucImpl.WriteRequestUsecaseDto{
		MetadataRepo:        &metadataRepo,
		NodeStorageRepo:     &nodeStorageRepo,
		INodeRepo:           &iNodeRepo,
		ServiceRegistry:     &serviceRegistry,
		DataNodeCache:       dataNodeCache,
		TransactionsRepo:    &transactionRepo,
		Mtx:                 dataNodeCachemtx,
		Cfg:                 cfg,
		NodeAllocator:       &nodeAllocatorSvc,
		DataNodeService:     &dataNodeSvc,
		RollbackService:     &rollbackSvc,
		TransactionInjector: transactionInjector,
	})
	cronUC := ucImpl.NewCronUsecase(&ucImpl.CronUsecaseDto{
		ServiceRegistry:  &serviceRegistry,
		DataNodeCache:    dataNodeCache,
		Mtx:              dataNodeCachemtx,
		TransactionsRepo: &transactionRepo,
		RollbackService:  &rollbackSvc,
	})

	// init delivery
	server := grpc.NewGrpcHandler(cfg, &writeUC)

	// spawn cron on another thread
	cron := time.NewTicker(5 * time.Second)
	defer cron.Stop()
	go func(ch <-chan time.Time) {
		for t := range ch {

			log.Printf("%s cron started\n", t.Local().String())

			ctx := context.Background()

			errGroup := &errgroup.Group{}

			// clean up expired transaction commit
			errGroup.Go(func() error {
				return cronUC.TransactionCleanUp(ctx)
			})

			// cache dataNode service entry registry
			errGroup.Go(func() error {
				return cronUC.SetDataNodeCache(ctx)
			})

			err := errGroup.Wait()
			log.Printf("err %s", err.Error())
		}
	}(cron.C)

	// serve
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.NameNodePort))
	log.Printf("gRPC server listening on %s", lis.Addr())

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
