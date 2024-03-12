package impl

import (
	"sync"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	pkgRepo "github.com/hdkef/hadoop/pkg/repository/transactionable"
	pkgSvc "github.com/hdkef/hadoop/pkg/services"
	pkgSvcImpl "github.com/hdkef/hadoop/pkg/services/impl"
	"github.com/hdkef/hadoop/services/nameNode/config"
	"github.com/hdkef/hadoop/services/nameNode/repository"
	"github.com/hdkef/hadoop/services/nameNode/service"
	svcImpl "github.com/hdkef/hadoop/services/nameNode/service/impl"
	"github.com/hdkef/hadoop/services/nameNode/usecase"
)

type WriteRequestUsecaseImpl struct {
	metadataRepo        repository.MetadataRepo
	nodeStorageRepo     repository.NodeStorageRepo
	iNodeRepo           repository.INodeRepo
	serviceRegistry     pkgSvc.ServiceRegistry
	dataNodeCache       map[string]*pkgEt.ServiceDiscovery
	transactionsRepo    repository.TransactionsRepo
	mtx                 *sync.Mutex
	cfg                 *config.Config
	nodeAllocator       service.NodeAllocator
	dataNodeService     service.DataNodeService
	rollbackService     service.RollbackService
	transactionInjector *pkgRepo.TransactionInjector
}

func NewWriteUsecase(cfg *config.Config, dataNodeCache map[string]*pkgEt.ServiceDiscovery, mtx *sync.Mutex) usecase.WriteRequestUsecase {
	return &WriteRequestUsecaseImpl{
		dataNodeCache:       dataNodeCache,
		mtx:                 mtx,
		nodeAllocator:       svcImpl.NewNodeAllocator(),
		dataNodeService:     svcImpl.NewDataNodeService(),
		serviceRegistry:     pkgSvcImpl.NewServiceRegistry(),
		cfg:                 cfg,
		transactionInjector: pkgRepo.NewTransactionInjector(nil),
	}
}
