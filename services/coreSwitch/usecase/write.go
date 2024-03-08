package usecase

import (
	"context"

	"github.com/hdkef/hadoop/services/coreSwitch/entity"
)

type WriteUsecase interface {
	Write(ctx context.Context, dto *entity.WriteDto, chProgress chan uint8) error
}
