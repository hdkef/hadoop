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
	tx, err := w.transactionsRepo.Get(ctx, txID, nil)
	if err != nil || tx == nil {
		return errors.New("transactions not found")
	}

	// begin db transactions
	transactionable, err := w.transactionInjector.Begin(ctx)
	if err != nil || tx == nil {
		return err
	}
	// create iNodeblockids

	iNode := &entity.INode{}
	iNode.SetID(tx.GetID())
	iNode.SetHash(tx.GetMetadata().GetHash())
	iNode.SetBlocks(tx.GetBlockTaret())

	err = w.iNodeRepo.Create(ctx, iNode, transactionable)
	if err != nil {
		transactionable.Rollback()
		return err
	}

	// create metadata

	err = w.metadataRepo.Touch(ctx, tx.GetMetadata(), transactionable)
	if err != nil {
		transactionable.Rollback()
		return err
	}

	// update transaction checkpoint
	err = w.transactionsRepo.Commit(ctx, tx.GetID(), transactionable)
	if err != nil {
		transactionable.Rollback()
		return err
	}

	// commit db transactions
	transactionable.Commit()

	return nil
}
