package transactionable

import (
	"context"
	"database/sql"

	"github.com/hdkef/hadoop/pkg/repository"
)

type TransactionInjector struct {
	db *sql.DB
}

// Begin implements repository.Transactionable.
func (t *TransactionInjector) Begin(ctx context.Context) (repository.Transactionable, error) {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Commit implements repository.Transactionable.
func (t *TransactionInjector) Commit(tx repository.Transactionable) error {
	return tx.Commit()
}

// Rollback implements repository.Transactionable.
func (t *TransactionInjector) Rollback(tx repository.Transactionable) error {
	return tx.Commit()
}

func NewTransactionInjector(db *sql.DB) *TransactionInjector {
	return &TransactionInjector{
		db: db,
	}
}
