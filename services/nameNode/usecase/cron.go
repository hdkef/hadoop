package usecase

import (
	"context"
	"sync"

	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type CronUsecase interface {
	TransactionCleanUp(ctx context.Context) error
	SetDataNodeCache(ctx context.Context, svd map[string]*entity.ServiceDiscovery, mtx *sync.Mutex) error
}
