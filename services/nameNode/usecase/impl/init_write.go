package impl

import (
	"sync"

	"github.com/hdkef/hadoop/services/nameNode/config"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/repository"
	"github.com/hdkef/hadoop/services/nameNode/service"
	svcImpl "github.com/hdkef/hadoop/services/nameNode/service/impl"
	"github.com/hdkef/hadoop/services/nameNode/usecase"
)

type WriteRequestUsecaseImpl struct {
	metadataRepo     repository.MetadataRepo
	nodeStorageRepo  repository.NodeStorageRepo
	iNodeRepo        repository.INodeRepo
	serviceRegistry  service.ServiceRegistry
	dataNodeCache    map[string]*entity.ServiceDiscovery
	transactionsRepo repository.TransactionsRepo
	mtx              *sync.Mutex
	cfg              *config.Config
	nodeAllocator    service.NodeAllocator
	dataNodeService  service.DataNodeService
}

func NewWriteUsecase(cfg *config.Config, dataNodeCache map[string]*entity.ServiceDiscovery, mtx *sync.Mutex) usecase.WriteRequestUsecase {
	return &WriteRequestUsecaseImpl{
		dataNodeCache:   dataNodeCache,
		mtx:             mtx,
		nodeAllocator:   svcImpl.NewNodeAllocator(),
		dataNodeService: svcImpl.NewDataNodeService(),
		serviceRegistry: svcImpl.NewServiceRegistry(cfg),
		cfg:             cfg,
	}
}
