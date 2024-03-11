package repository

import (
	"context"
	"time"

	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type JobQueueRepo interface {
	Set(ctx context.Context, et *entity.JobQueue, ttl *time.Duration) error
	Get(ctx context.Context, key string) (*entity.JobQueue, error)
	Del(ctx context.Context, key string)
}
