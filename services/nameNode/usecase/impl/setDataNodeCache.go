package impl

import (
	"context"

	"github.com/hdkef/hadoop/pkg/logger"
)

// SetDataNodeCache implements usecase.CronUsecase.
func (c *CronUsecase) SetDataNodeCache(ctx context.Context) error {

	// get dataNode via service discovery
	svds, err := c.serviceRegistry.GetAll(ctx, "dataNode", "")
	if err != nil {
		logger.LogError(err)
		return err
	}
	c.mtx.Lock()
	defer c.mtx.Unlock()
	// delete old registry
	for key := range c.dataNodeCache {
		delete(c.dataNodeCache, key)
	}
	// set new registry
	for _, v := range svds {
		c.dataNodeCache[v.GetID()] = v
	}

	return nil
}
