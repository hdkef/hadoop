package impl

import (
	"sync"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	pkgRepo "github.com/hdkef/hadoop/pkg/repository/transactionable"
	pkgSvc "github.com/hdkef/hadoop/pkg/services"
	"github.com/hdkef/hadoop/services/nameNode/config"
	"github.com/hdkef/hadoop/services/nameNode/repository"
	"github.com/hdkef/hadoop/services/nameNode/service"
	"github.com/hdkef/hadoop/services/nameNode/usecase"
)

type WriteRequestUsecaseDto struct {
	MetadataRepo        *repository.MetadataRepo
	NodeStorageRepo     *repository.NodeStorageRepo
	INodeRepo           *repository.INodeRepo
	ServiceRegistry     *pkgSvc.ServiceRegistry
	DataNodeCache       map[string]*pkgEt.ServiceDiscovery
	TransactionsRepo    *repository.TransactionsRepo
	Mtx                 *sync.Mutex
	Cfg                 *config.Config
	NodeAllocator       *service.NodeAllocator
	DataNodeService     *service.DataNodeService
	RollbackService     *service.RollbackService
	TransactionInjector *pkgRepo.TransactionInjector
}

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

func NewWriteUsecase(dto *WriteRequestUsecaseDto) usecase.WriteRequestUsecase {

	if dto.MetadataRepo == nil {
		panic("metadataRepo nil")
	}

	if dto.NodeStorageRepo == nil {
		panic("nodeStorageRepo nil")
	}

	if dto.INodeRepo == nil {
		panic("iNodeRepo nil")
	}

	if dto.ServiceRegistry == nil {
		panic("serviceRegistry nil")
	}

	if dto.TransactionsRepo == nil {
		panic("transactionsRepo nil")
	}

	if dto.Mtx == nil {
		panic("mtx nil")
	}

	if dto.Cfg == nil {
		panic("cfg nil")
	}

	if dto.NodeAllocator == nil {
		panic("nodeAllocator nil")
	}

	if dto.DataNodeService == nil {
		panic("dataNodeService nil")
	}

	if dto.RollbackService == nil {
		panic("rollbackService nil")
	}

	if dto.TransactionInjector == nil {
		panic("transactionInjector nil")
	}

	return &WriteRequestUsecaseImpl{
		metadataRepo:        *dto.MetadataRepo,
		nodeStorageRepo:     *dto.NodeStorageRepo,
		iNodeRepo:           *dto.INodeRepo,
		serviceRegistry:     *dto.ServiceRegistry,
		transactionsRepo:    *dto.TransactionsRepo,
		mtx:                 dto.Mtx,
		cfg:                 dto.Cfg,
		nodeAllocator:       *dto.NodeAllocator,
		dataNodeService:     *dto.DataNodeService,
		rollbackService:     *dto.RollbackService,
		transactionInjector: dto.TransactionInjector,
	}
}
