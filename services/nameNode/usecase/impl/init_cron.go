package impl

import (
	"sync"

	"github.com/hdkef/hadoop/services/nameNode/config"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/service"
	svcImpl "github.com/hdkef/hadoop/services/nameNode/service/impl"
	"github.com/hdkef/hadoop/services/nameNode/usecase"
)

type CronUsecase struct {
	serviceRegistry service.ServiceRegistry
	dataNodeCache   map[string]*entity.ServiceDiscovery
	mtx             *sync.Mutex
}

func NewCronUsecase(cfg *config.Config, dataNodeCache map[string]*entity.ServiceDiscovery, mtx *sync.Mutex) usecase.CronUsecase {
	return &CronUsecase{
		dataNodeCache:   dataNodeCache,
		mtx:             mtx,
		serviceRegistry: svcImpl.NewServiceRegistry(cfg),
	}
}
