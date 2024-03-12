package postgres

import (
	"context"
	"database/sql"

	pkgRepo "github.com/hdkef/hadoop/pkg/repository"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/repository"
)

type INodeRepo struct {
	db *sql.DB
}

// Create implements repository.INodeRepo.
func (i *INodeRepo) Create(ctx context.Context, inode *entity.INode, tx pkgRepo.Transactionable) error {
	panic("unimplemented")
}

// Delete implements repository.INodeRepo.
func (i *INodeRepo) Delete(ctx context.Context, inodeID string, tx pkgRepo.Transactionable) error {
	panic("unimplemented")
}

// Get implements repository.INodeRepo.
func (i *INodeRepo) Get(ctx context.Context, inodeID string, tx pkgRepo.Transactionable) (*entity.INode, error) {
	panic("unimplemented")
}

func NewINodeRepo(db *sql.DB) repository.INodeRepo {
	return &INodeRepo{
		db: db,
	}
}
