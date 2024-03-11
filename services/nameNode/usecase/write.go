package usecase

import (
	"context"

	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type WriteRequestUsecase interface {
	CreateRequest(ctx context.Context, dto *entity.CreateReqDto) error
	CommitCreateRequest(ctx context.Context) error
	CheckDataNode(ctx context.Context) error
}
