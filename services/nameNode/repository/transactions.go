package repository

import (
	"context"

	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type TransactionsRepo interface {
	Add(ctx context.Context, et *entity.Transactions) error
	Commit(ctx context.Context)
	Get(ctx context.Context)
}
