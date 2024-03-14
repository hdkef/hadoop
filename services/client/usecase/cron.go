package usecase

import "context"

type CronUsecase interface {
	SetNameNodeCache(ctx context.Context) error
}
