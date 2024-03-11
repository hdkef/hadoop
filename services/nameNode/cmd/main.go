package cmd

import (
	"time"
)

func main() {

	// spawn cron on another thread
	cron := time.NewTicker(1 * time.Second)
	defer cron.Stop()
	go func(ch <-chan time.Time) {
		for t := range ch {
			// clean up expired transaction commit

			// cache dataNode service entry registry
		}
	}(cron.C)
}
