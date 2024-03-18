package impl

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hdkef/hadoop/pkg/logger"
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
	if err != nil {
		logger.LogError(err)
		return err
	}

	// create metadata

	err = w.metadataRepo.Touch(ctx, tx.GetMetadata(), transactionable)
	if err != nil {
		logger.LogError(err)
		transactionable.Tx.Rollback()
		return err
	}

	// create iNodeblockids

	iNode := &entity.INode{}
	iNode.SetID(tx.GetMetadata().GetINodeID())
	iNode.SetBlocks(tx.GetBlockTaret())
	iNode.SetAllBlockIds(tx.GetMetadata().GetAllBlockIds())

	err = w.iNodeRepo.Create(ctx, iNode, transactionable)
	if err != nil {
		logger.LogError(err)
		transactionable.Tx.Rollback()
		return err
	}

	// update transaction checkpoint
	err = w.transactionsRepo.Commit(ctx, tx.GetID(), transactionable)
	if err != nil {
		logger.LogError(err)
		transactionable.Tx.Rollback()
		return err
	}

	// commit db transactions
	err = transactionable.Tx.Commit()
	if err != nil {
		logger.LogError(err)
		transactionable.Tx.Rollback()
		return err
	}
	return nil
}
