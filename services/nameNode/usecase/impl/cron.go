package impl

import (
	"context"
	"sync"

	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/usecase"
)

type CronUsecase struct {
}

// SetDataNodeCache implements usecase.CronUsecase.
func (c *CronUsecase) SetDataNodeCache(ctx context.Context, svd map[string]*entity.ServiceDiscovery, mtx *sync.Mutex) error {
	panic("unimplemented")
}

// TransactionCleanUp implements usecase.CronUsecase.
func (c *CronUsecase) TransactionCleanUp(ctx context.Context) error {
	panic("unimplemented")
}

func NewCronUsecase() usecase.CronUsecase {
	return &CronUsecase{}
}
