package impl

import "context"

// SetNameNodeCache implements usecase.CronUsecase.
func (c *CronUsecase) SetNameNodeCache(ctx context.Context) error {
	// get dataNode via service discovery
	svds, err := c.serviceRegistry.GetAll(ctx, "dataNode", "")
	if err != nil {
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
