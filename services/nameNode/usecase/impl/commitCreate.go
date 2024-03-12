package impl

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hdkef/hadoop/services/nameNode/entity"
)

// CommitCreateRequest implements usecase.WriteRequestUsecase.
func (w *WriteRequestUsecaseImpl) CommitTransactions(ctx context.Context, txID uuid.UUID) error {

	// get transactions
	tx, err := w.transactionsRepo.Get(ctx, txID)
	if err != nil || tx == nil {
		return errors.New("transactions not found")
	}

	// create iNodeblockids

	iNode := &entity.INode{}
	iNode.SetID(tx.GetID())
	iNode.SetHash(tx.GetMetadata().GetHash())
	iNode.SetBlocks(tx.GetBlockTaret())

	err = w.iNodeRepo.Create(ctx, iNode)
	if err != nil {
		return err
	}

	// update transaction checkpoint
	err = w.transactionsRepo.Commit(ctx, tx.GetID())
	if err != nil {
		return err
	}

	return nil
}
