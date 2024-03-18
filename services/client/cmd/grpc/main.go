package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/hdkef/hadoop/services/client/config"
	"github.com/hdkef/hadoop/services/client/delivery/grpc"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	pkgSvc "github.com/hdkef/hadoop/pkg/services/impl"
	svcImpl "github.com/hdkef/hadoop/services/client/service/impl"
	ucImpl "github.com/hdkef/hadoop/services/client/usecase/impl"
)

func main() {

	// init config
	cfg := config.NewConfigServer()

	// init service

	nameNodeCache := make(map[int]*pkgEt.ServiceDiscovery)
	nameNodeCacheMtx := &sync.Mutex{}

	serviceRegistry := pkgSvc.NewServiceRegistry(cfg.ServiceRegistryConfig)
	dataNodeService := svcImpl.NewDataNodeService()
	nameNodeService := svcImpl.NewNameNodeService(&svcImpl.NameNodeServiceDto{
		ServiceRegistry: &serviceRegistry,
		NameNodeCache:   nameNodeCache,
		Mtx:             nameNodeCacheMtx,
	})

	// init usecase
	writeUC := ucImpl.NewWriteUsecase(&dataNodeService, &nameNodeService)
	cronUC := ucImpl.NewCronUsecase(nameNodeCache, nameNodeCacheMtx, serviceRegistry)
	cron := time.NewTicker(5 * time.Second)
	defer cron.Stop()
	go func(ch <-chan time.Time) {
		for t := range ch {

			log.Printf("%s cron started\n", t.Local().String())

			ctx := context.Background()

			// cache dataNode service entry registry
			err := cronUC.SetNameNodeCache(ctx)

			if err != nil {
				log.Printf("err %s", err.Error())
			}
		}
	}(cron.C)

	// spin up tcp listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.ClientPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("gRPC server listening on %s", lis.Addr())

	// spin up grpc server
	sv := grpc.NewGrpcHandler(cfg, writeUC)
	if err := sv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
