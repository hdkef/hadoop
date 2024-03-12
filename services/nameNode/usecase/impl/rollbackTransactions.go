package impl

import (
	"context"

	"github.com/google/uuid"
)

// RollbackTransactions implements usecase.WriteRequestUsecase.
func (w *WriteRequestUsecaseImpl) RollbackTransactions(ctx context.Context, transactionsID uuid.UUID) error {
	tx, err := w.transactionsRepo.Get(ctx, transactionsID)
	if err != nil {
		return err
	}
	return w.rollbackService.Rollback(ctx, tx)
}
