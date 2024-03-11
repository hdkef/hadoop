package usecase

import (
	"context"
)

type CronUsecase interface {
	TransactionCleanUp(ctx context.Context) error
	SetDataNodeCache(ctx context.Context) error
}
