package usecase

import (
	"context"

	"github.com/hdkef/hadoop/services/client/entity"
)

type WriteUsecase interface {
	Write(ctx context.Context, dto entity.WriteRequestDto) error
}
