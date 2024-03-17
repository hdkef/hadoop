package impl

import (
	"sync"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	pkgSvc "github.com/hdkef/hadoop/pkg/services"

	"github.com/hdkef/hadoop/services/client/usecase"
)

type CronUsecase struct {
	nameNodeCache   map[int]*pkgEt.ServiceDiscovery
	mtx             *sync.Mutex
	serviceRegistry pkgSvc.ServiceRegistry
}

func NewCronUsecase(nc map[int]*pkgEt.ServiceDiscovery, mtx *sync.Mutex, serviceRegistry pkgSvc.ServiceRegistry) usecase.CronUsecase {
	return &CronUsecase{
		nameNodeCache:   nc,
		mtx:             mtx,
		serviceRegistry: serviceRegistry,
	}
}
