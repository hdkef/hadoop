package repository

import "context"

type Transactions interface {
	Add(ctx context.Context) error
	Commit(ctx context.Context)
	Get(ctx context.Context)
}
