package impl

import (
	"context"

	"github.com/hdkef/hadoop/pkg/logger"
)

// SetNameNodeCache implements usecase.CronUsecase.
func (c *CronUsecase) SetNameNodeCache(ctx context.Context) error {
	// get dataNode via service discovery
	svds, err := c.serviceRegistry.GetAll(ctx, "nameNode", "")
	if err != nil {
		logger.LogError(err)
		return err
	}
	c.mtx.Lock()
	defer c.mtx.Unlock()
	// delete old registry
	for key := range c.nameNodeCache {
		delete(c.nameNodeCache, key)
	}
	// set new registry
	for i, v := range svds {
		c.nameNodeCache[i] = v
	}

	return nil
}
