package repository

import (
	"context"

	"github.com/google/uuid"
	pkgRepoTr "github.com/hdkef/hadoop/pkg/repository/transactionable"
	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type TransactionsRepo interface {
	Add(ctx context.Context, et *entity.Transactions, tx *pkgRepoTr.Transactionable) (uuid.UUID, error)
	Commit(ctx context.Context, transactionID uuid.UUID, tx *pkgRepoTr.Transactionable) error
	Get(ctx context.Context, transactionID uuid.UUID, tx *pkgRepoTr.Transactionable) (*entity.Transactions, error)
	GetOneExpired(ctx context.Context, tx *pkgRepoTr.Transactionable) (*entity.Transactions, error)
	RolledBack(ctx context.Context, transactionID uuid.UUID, tx *pkgRepoTr.Transactionable) error
	IncrementRollbackRetries(ctx context.Context, transactionID uuid.UUID, tx *pkgRepoTr.Transactionable) error
}
