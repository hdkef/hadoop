package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	pkgRepo "github.com/hdkef/hadoop/pkg/repository"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/repository"
)

type TransactionsRepo struct {
	db *sql.DB
}

// Add implements repository.TransactionsRepo.
func (t *TransactionsRepo) Add(ctx context.Context, et *entity.Transactions, tx pkgRepo.Transactionable) error {
	panic("unimplemented")
}

// Commit implements repository.TransactionsRepo.
func (t *TransactionsRepo) Commit(ctx context.Context, transactionID uuid.UUID, tx pkgRepo.Transactionable) error {
	panic("unimplemented")
}

// Get implements repository.TransactionsRepo.
func (t *TransactionsRepo) Get(ctx context.Context, transactionID uuid.UUID, tx pkgRepo.Transactionable) (*entity.Transactions, error) {
	panic("unimplemented")
}

// GetOneExpired implements repository.TransactionsRepo.
func (t *TransactionsRepo) GetOneExpired(ctx context.Context, tx pkgRepo.Transactionable) (*entity.Transactions, error) {
	panic("unimplemented")
}

// RolledBack implements repository.TransactionsRepo.
func (t *TransactionsRepo) RolledBack(ctx context.Context, transactionID uuid.UUID, tx pkgRepo.Transactionable) error {
	panic("unimplemented")
}

func NewTransactionsRepo(db *sql.DB) repository.TransactionsRepo {
	return &TransactionsRepo{
		db: db,
	}
}
