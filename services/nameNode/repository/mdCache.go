package repository

import (
	"context"

	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type MDCacheRepo interface {
	GetMetadata(ctx context.Context, et *entity.Metadata) error
	SetMetadata(ctx context.Context, et *entity.Metadata) error
}
