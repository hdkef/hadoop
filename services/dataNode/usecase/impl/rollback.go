package impl

import (
	"context"

	"github.com/hdkef/hadoop/services/dataNode/entity"
)

// RollBack implements usecase.WriteUsecase.
func (w *WriteUsecaseImpl) RollBack(ctx context.Context, dto *entity.RollbackDto) error {
	panic("unimplemented")
}
