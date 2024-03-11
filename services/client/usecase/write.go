package usecase

import (
	"context"

	"github.com/hdkef/hadoop/services/client/entity"
)

type WriteUsecase interface {
	Create(ctx context.Context, dto *entity.CreateDto, chProgress chan entity.CreateStreamRes)
}
