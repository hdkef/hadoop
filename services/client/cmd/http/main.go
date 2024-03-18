package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/hdkef/hadoop/services/client/config"
	httpDelivery "github.com/hdkef/hadoop/services/client/delivery/http"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	pkgSvc "github.com/hdkef/hadoop/pkg/services/impl"
	svcImpl "github.com/hdkef/hadoop/services/client/service/impl"
	ucImpl "github.com/hdkef/hadoop/services/client/usecase/impl"
)

func main() {
	// init config
	cfg := config.NewConfigServer()

	// init service

	nameNodeCache := make(map[int]*pkgEt.ServiceDiscovery)
	nameNodeCacheMtx := &sync.Mutex{}

	serviceRegistry := pkgSvc.NewServiceRegistry(cfg.ServiceRegistryConfig)
	dataNodeService := svcImpl.NewDataNodeService()
	nameNodeService := svcImpl.NewNameNodeService(&svcImpl.NameNodeServiceDto{
		ServiceRegistry: &serviceRegistry,
		NameNodeCache:   nameNodeCache,
		Mtx:             nameNodeCacheMtx,
	})

	// init usecase
	writeUC := ucImpl.NewWriteUsecase(&dataNodeService, &nameNodeService)
	cronUC := ucImpl.NewCronUsecase(nameNodeCache, nameNodeCacheMtx, serviceRegistry)
	cron := time.NewTicker(60 * time.Second)
	defer cron.Stop()
	go func(ch <-chan time.Time) {
		for t := range ch {

			log.Printf("%s cron started\n", t.Local().String())

			ctx := context.Background()

			// cache dataNode service entry registry
			err := cronUC.SetNameNodeCache(ctx)

			if err != nil {
				log.Printf("err %s", err.Error())
			}
		}
	}(cron.C)

	// spin up grpc server
	sv := httpDelivery.NewHTTPHandler(cfg, writeUC)

	log.Printf("http server listening on %d", cfg.ClientPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.ClientPort), sv); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
