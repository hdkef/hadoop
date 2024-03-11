package impl

import (
	"context"
)

// SetDataNodeCache implements usecase.CronUsecase.
func (c *CronUsecase) SetDataNodeCache(ctx context.Context) error {

	// get dataNode via service discovery
	svds, err := c.serviceRegistry.GetAll(ctx, "dataNode", "")
	if err != nil {
		return err
	}
	c.mtx.Lock()
	defer c.mtx.Unlock()
	for _, v := range svds {
		c.dataNodeCache[v.GetID()] = v
	}

	return nil
}
