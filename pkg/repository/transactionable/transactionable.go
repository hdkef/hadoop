package transactionable

import (
	"context"
	"database/sql"
)

type Transactionable struct {
	Tx *sql.Tx
}

type TransactionInjector struct {
	db *sql.DB
}

// Begin implements Transactionable.
func (t *TransactionInjector) Begin(ctx context.Context) (*Transactionable, error) {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &Transactionable{
		Tx: tx,
	}, nil
}

// Commit implements Transactionable.
func (t *TransactionInjector) Commit(tx *Transactionable) error {
	return tx.Tx.Commit()
}

// Rollback implements Transactionable.
func (t *TransactionInjector) Rollback(tx *Transactionable) error {
	return tx.Tx.Commit()
}

func NewTransactionInjector(db *sql.DB) *TransactionInjector {
	return &TransactionInjector{
		db: db,
	}
}
