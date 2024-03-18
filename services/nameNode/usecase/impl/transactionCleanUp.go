package impl

import (
	"context"

	"github.com/hdkef/hadoop/pkg/logger"
)

// TransactionCleanUp implements usecase.CronUsecase.
func (c *CronUsecase) TransactionCleanUp(ctx context.Context) error {

	// get one expired transactions
	tx, err := c.transactionsRepo.GetOneExpired(ctx, nil)
	if err != nil {
		logger.LogError(err)
		return err
	}

	if tx == nil {
		return nil
	}

	// execute rollback

	err = c.rollbackService.Rollback(ctx, tx)
	if err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
