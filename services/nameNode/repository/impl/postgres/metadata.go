package postgres

import (
	"context"
	"database/sql"

	pkgRepo "github.com/hdkef/hadoop/pkg/repository"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/repository"
)

type MetadataRepo struct {
	db *sql.DB
}

// CheckPath implements repository.MetadataRepo.
func (m *MetadataRepo) CheckPath(ctx context.Context, path string, tx pkgRepo.Transactionable) bool {
	panic("unimplemented")
}

// Delete implements repository.MetadataRepo.
func (m *MetadataRepo) Delete(ctx context.Context, metadata *entity.Metadata, tx pkgRepo.Transactionable) error {
	panic("unimplemented")
}

// Get implements repository.MetadataRepo.
func (m *MetadataRepo) Get(ctx context.Context, path string, tx pkgRepo.Transactionable) (*entity.Metadata, error) {
	panic("unimplemented")
}

// Touch implements repository.MetadataRepo.
func (m *MetadataRepo) Touch(ctx context.Context, et *entity.Metadata, tx pkgRepo.Transactionable) error {
	panic("unimplemented")
}

func NewMetadataRepo(db *sql.DB) repository.MetadataRepo {
	return &MetadataRepo{db: db}
}
