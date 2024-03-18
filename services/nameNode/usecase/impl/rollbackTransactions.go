package impl

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hdkef/hadoop/pkg/logger"
)

// RollbackTransactions implements usecase.WriteRequestUsecase.
func (w *WriteRequestUsecaseImpl) RollbackTransactions(ctx context.Context, transactionsID uuid.UUID) error {
	tx, err := w.transactionsRepo.Get(ctx, transactionsID, nil)
	if err != nil || tx == nil {
		return errors.New("transactions not found")
	}
	err = w.rollbackService.Rollback(ctx, tx)
	if err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
