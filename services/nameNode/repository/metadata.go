package repository

import (
	"context"

	pkgRepoTr "github.com/hdkef/hadoop/pkg/repository/transactionable"
	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type MetadataRepo interface {
	CheckPath(ctx context.Context, path string, tx *pkgRepoTr.Transactionable) bool
	Touch(ctx context.Context, et *entity.Metadata, tx *pkgRepoTr.Transactionable) error
	Get(ctx context.Context, path string, tx *pkgRepoTr.Transactionable) (*entity.Metadata, error)
	Delete(ctx context.Context, metadata *entity.Metadata, tx *pkgRepoTr.Transactionable) error
}
