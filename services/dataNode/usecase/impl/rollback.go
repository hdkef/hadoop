package impl

import (
	"context"

	"github.com/hdkef/hadoop/services/dataNode/entity"
)

// RollBack implements usecase.WriteUsecase.
func (w *WriteUsecaseImpl) RollBack(ctx context.Context, dto *entity.RollbackDto) error {

	// remove from storage
	iNodeBlockID := &entity.INodeBlockID{}
	iNodeBlockID.SetINodeID(dto.GetINodeID())
	iNodeBlockID.SetBlockID(dto.GetBlockID())

	return iNodeBlockID.Remove(w.cfg.StorageRoot)
}
