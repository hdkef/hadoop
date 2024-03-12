package impl

import (
	"context"
)

// TransactionCleanUp implements usecase.CronUsecase.
func (c *CronUsecase) TransactionCleanUp(ctx context.Context) error {

	// get one expired transactions
	tx, err := c.transactionsRepo.GetOneExpired(ctx)
	if err != nil {
		return err
	}

	if tx == nil {
		return nil
	}

	// execute rollback

	return c.rollbackService.Rollback(ctx, tx)
}
