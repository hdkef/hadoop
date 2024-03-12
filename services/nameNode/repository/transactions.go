package repository

import (
	"context"

	"github.com/google/uuid"
	pkgRepo "github.com/hdkef/hadoop/pkg/repository"
	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type TransactionsRepo interface {
	Add(ctx context.Context, et *entity.Transactions, tx pkgRepo.Transactionable) error
	Commit(ctx context.Context, transactionID uuid.UUID, tx pkgRepo.Transactionable) error
	Get(ctx context.Context, transactionID uuid.UUID, tx pkgRepo.Transactionable) (*entity.Transactions, error)
	GetOneExpired(ctx context.Context, tx pkgRepo.Transactionable) (*entity.Transactions, error)
	RolledBack(ctx context.Context, transactionID uuid.UUID, tx pkgRepo.Transactionable) error
}
