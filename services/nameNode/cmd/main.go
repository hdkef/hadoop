package cmd

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/hdkef/hadoop/services/nameNode/entity"
	ucImpl "github.com/hdkef/hadoop/services/nameNode/usecase/impl"
	"golang.org/x/sync/errgroup"
)

func main() {

	dataNodeCache := make(map[string]*entity.ServiceDiscovery)
	mtx := &sync.Mutex{}
	cronUC := ucImpl.NewCronUsecase()

	// spawn cron on another thread
	cron := time.NewTicker(1 * time.Second)
	defer cron.Stop()
	go func(ch <-chan time.Time) {
		for t := range ch {

			log.Printf("%s cron started\n", t.Local().String())

			ctx := context.Background()

			errGroup := &errgroup.Group{}

			// clean up expired transaction commit
			errGroup.Go(func() error {
				return cronUC.TransactionCleanUp(ctx)
			})

			// cache dataNode service entry registry
			errGroup.Go(func() error {
				return cronUC.SetDataNodeCache(ctx, dataNodeCache, mtx)
			})

			err := errGroup.Wait()
			log.Printf("err %s", err.Error())
		}
	}(cron.C)
}
