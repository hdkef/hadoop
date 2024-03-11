package usecase

import (
	"context"

	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type CronUsecase interface {
	TransactionCleanUp(ctx context.Context)
	SetDataNodeCache(ctx context.Context, svd map[string]*entity.ServiceDiscovery)
}
