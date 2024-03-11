package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type TransactionsRepo interface {
	Add(ctx context.Context, et *entity.Transactions) error
	Commit(ctx context.Context, transactionID uuid.UUID) error
	Get(ctx context.Context, transactionID uuid.UUID) (*entity.Transactions, error)
}
