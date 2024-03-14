package impl

import (
	"sync"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	pkgSvc "github.com/hdkef/hadoop/pkg/services"
	"github.com/hdkef/hadoop/services/nameNode/repository"
	"github.com/hdkef/hadoop/services/nameNode/service"
	"github.com/hdkef/hadoop/services/nameNode/usecase"
)

type CronUsecaseDto struct {
	ServiceRegistry  *pkgSvc.ServiceRegistry
	DataNodeCache    map[string]*pkgEt.ServiceDiscovery
	Mtx              *sync.Mutex
	TransactionsRepo *repository.TransactionsRepo
	RollbackService  *service.RollbackService
}

type CronUsecase struct {
	serviceRegistry  pkgSvc.ServiceRegistry
	dataNodeCache    map[string]*pkgEt.ServiceDiscovery
	mtx              *sync.Mutex
	transactionsRepo repository.TransactionsRepo
	rollbackService  service.RollbackService
}

func NewCronUsecase(dto *CronUsecaseDto) usecase.CronUsecase {

	if dto.DataNodeCache == nil {
		panic("dataNodeCache is nil")
	}

	if dto.Mtx == nil {
		panic("mtx is nil")
	}

	if dto.ServiceRegistry == nil {
		panic("serviceRegistry is nil")
	}

	if dto.TransactionsRepo == nil {
		panic("transactionsRepo is nil")
	}

	return &CronUsecase{
		dataNodeCache:    dto.DataNodeCache,
		mtx:              dto.Mtx,
		serviceRegistry:  *dto.ServiceRegistry,
		transactionsRepo: *dto.TransactionsRepo,
	}
}
