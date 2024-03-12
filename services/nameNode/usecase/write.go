package usecase

import (
	"context"

	"github.com/google/uuid"
	pkgEt "github.com/hdkef/hadoop/pkg/entity"
)

type WriteRequestUsecase interface {
	CreateRequest(ctx context.Context, dto *pkgEt.CreateReqDto) ([]*pkgEt.QueryNodeTarget, error)
	CommitTransactions(ctx context.Context, transactionsID uuid.UUID) error
	RollbackTransactions(ctx context.Context, transactionsID uuid.UUID) error
}
