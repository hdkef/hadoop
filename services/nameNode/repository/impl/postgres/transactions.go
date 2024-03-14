package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	pkgRepoTr "github.com/hdkef/hadoop/pkg/repository/transactionable"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/repository"
)

type TransactionsRepo struct {
	db *sql.DB
}

// Add implements repository.TransactionsRepo.
func (t *TransactionsRepo) Add(ctx context.Context, et *entity.Transactions, tx *pkgRepoTr.Transactionable) error {
	panic("unimplemented")
}

// Commit implements repository.TransactionsRepo.
func (t *TransactionsRepo) Commit(ctx context.Context, transactionID uuid.UUID, tx *pkgRepoTr.Transactionable) error {
	panic("unimplemented")
}

// Get implements repository.TransactionsRepo.
func (t *TransactionsRepo) Get(ctx context.Context, transactionID uuid.UUID, tx *pkgRepoTr.Transactionable) (*entity.Transactions, error) {
	panic("unimplemented")
}

// GetOneExpired implements repository.TransactionsRepo.
func (t *TransactionsRepo) GetOneExpired(ctx context.Context, tx *pkgRepoTr.Transactionable) (*entity.Transactions, error) {
	panic("unimplemented")
}

// RolledBack implements repository.TransactionsRepo.
func (t *TransactionsRepo) RolledBack(ctx context.Context, transactionID uuid.UUID, tx *pkgRepoTr.Transactionable) error {
	panic("unimplemented")
}

func NewTransactionsRepo(db *sql.DB) repository.TransactionsRepo {
	return &TransactionsRepo{
		db: db,
	}
}
