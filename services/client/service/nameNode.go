package service

import (
	"context"

	"github.com/google/uuid"
	pkgEt "github.com/hdkef/hadoop/pkg/entity"
)

type NameNodeService interface {
	QueryNodeTarget(ctx context.Context, dto *pkgEt.CreateReqDto) (*pkgEt.QueryNodeTarget, error)
	CommitTransaction(ctx context.Context, transactionsID uuid.UUID, isSuccess bool) error
}
