package repository

import (
	"context"

	pkgRepo "github.com/hdkef/hadoop/pkg/repository"
	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type MetadataRepo interface {
	CheckPath(ctx context.Context, path string, tx pkgRepo.Transactionable) bool
	Touch(ctx context.Context, et *entity.Metadata, tx pkgRepo.Transactionable) error
	Get(ctx context.Context, path string, tx pkgRepo.Transactionable) (*entity.Metadata, error)
	Delete(ctx context.Context, metadata *entity.Metadata, tx pkgRepo.Transactionable) error
}
