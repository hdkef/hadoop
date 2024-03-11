package repository

import (
	"context"

	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type MetadataRepo interface {
	CheckPath(ctx context.Context, path string) bool
	Mkdir(ctx context.Context, path string) error
	Touch(ctx context.Context, path string, iNodeID string) error
	Get(ctx context.Context, path string) (*entity.Metadata, error)
	Delete(ctx context.Context, path string) error
}
