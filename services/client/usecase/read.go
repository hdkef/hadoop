package usecase

import (
	"context"

	"github.com/hdkef/hadoop/services/client/entity"
)

type ReadUsecase interface {
	Read(ctx context.Context, dto entity.ReadRequestDto) ([]byte, error)
}
