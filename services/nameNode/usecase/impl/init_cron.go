package impl

import (
	"sync"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	pkgSvc "github.com/hdkef/hadoop/pkg/services"
	pkgSvcImpl "github.com/hdkef/hadoop/pkg/services/impl"
	"github.com/hdkef/hadoop/services/nameNode/config"
	"github.com/hdkef/hadoop/services/nameNode/repository"
	svcImpl "github.com/hdkef/hadoop/services/nameNode/service/impl"
	"github.com/hdkef/hadoop/services/nameNode/usecase"
)

type CronUsecase struct {
	serviceRegistry  pkgSvc.ServiceRegistry
	dataNodeCache    map[string]*pkgEt.ServiceDiscovery
	mtx              *sync.Mutex
	transactionsRepo repository.TransactionsRepo
	metadataRepo     repository.MetadataRepo
	dataNodeService  svcImpl.DataNodeService
}

func NewCronUsecase(cfg *config.Config, dataNodeCache map[string]*pkgEt.ServiceDiscovery, mtx *sync.Mutex, transactionsRepo repository.TransactionsRepo) usecase.CronUsecase {
	return &CronUsecase{
		dataNodeCache:    dataNodeCache,
		mtx:              mtx,
		serviceRegistry:  pkgSvcImpl.NewServiceRegistry(),
		transactionsRepo: transactionsRepo,
	}
}
