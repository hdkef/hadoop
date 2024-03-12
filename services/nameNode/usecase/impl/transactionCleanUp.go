package impl

import (
	"context"

	"github.com/hdkef/hadoop/services/nameNode/entity"
	"golang.org/x/sync/errgroup"
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

	errGroup := &errgroup.Group{}

	// execute rollback

	// remove metadata
	errGroup.Go(func() error {
		return c.metadataRepo.Delete(ctx, tx.GetMetadata())
	})

	// remove files in dataNode
	blocks := tx.GetBlockTaret()
	for _, v := range blocks {
		for _, k := range v.NodeIDs {

			c.mtx.Lock()
			dto := &entity.RollbackDto{}
			dto.SetBlockID(v.ID)
			dto.SetINodeID(tx.GetMetadata().GetINodeID())
			dto.SetNodeAddress(c.dataNodeCache[k].GetAddress())
			dto.SetNodePort(c.dataNodeCache[k].GetPort())
			dto.SetNodeID(k)
			c.mtx.Unlock()

			errGroup.Go(func() error {
				return c.dataNodeService.Rollback(ctx, dto)
			})
		}
	}

	err = errGroup.Wait()
	if err != nil {
		return err
	}

	// rolled back transaction
	return c.transactionsRepo.RolledBack(ctx, tx.GetID())
}
