package dragonfly

import (
	"context"
	"time"

	pkgRepo "github.com/hdkef/hadoop/pkg/repository"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/repository"
)

type JobQueueRepo struct {
	client pkgRepo.KeyValueRepository
}

// Del implements repository.JobQueueRepo.
func (j *JobQueueRepo) Del(ctx context.Context, key string) {
	panic("unimplemented")
}

// Get implements repository.JobQueueRepo.
func (j *JobQueueRepo) Get(ctx context.Context, key string) (*entity.JobQueue, error) {
	panic("unimplemented")
}

// Set implements repository.JobQueueRepo.
func (j *JobQueueRepo) Set(ctx context.Context, et *entity.JobQueue, ttl *time.Duration) error {
	panic("unimplemented")
}

func NewJobQueueRepo(client pkgRepo.KeyValueRepository) repository.JobQueueRepo {
	return &JobQueueRepo{
		client: client,
	}
}
