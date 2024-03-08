package usecase

import (
	"context"

	"github.com/hdkef/hadoop/services/dataNode/entity"
)

type WriteUsecase interface {
	Write(ctx context.Context, dto *entity.WriteDto) error
}
