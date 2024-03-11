package usecase

import (
	"context"

	"github.com/google/uuid"
	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type WriteRequestUsecase interface {
	CreateRequest(ctx context.Context, dto *entity.CreateReqDto) ([]*pkgEt.QueryNodeTarget, error)
	CommitCreateRequest(ctx context.Context, transactionsID uuid.UUID, hash string) error
}
