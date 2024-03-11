package usecase

import (
	"context"

	"github.com/hdkef/hadoop/services/dataNode/entity"
)

type WriteUsecase interface {
	Create(ctx context.Context, dto *entity.CreateDto) error
	RollBack(ctx context.Context, dto *entity.RollbackDto) error
}
