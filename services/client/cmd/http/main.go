package main

// import (
// 	"sync"

// 	"github.com/hdkef/hadoop/services/client/config"

// 	pkgEt "github.com/hdkef/hadoop/pkg/entity"
// 	pkgSvc "github.com/hdkef/hadoop/pkg/services/impl"
// 	svcImpl "github.com/hdkef/hadoop/services/client/service/impl"
// 	ucImpl "github.com/hdkef/hadoop/services/client/usecase/impl"
// )

// func main() {
// 	// init config
// 	cfg := config.NewConfigServer()

// 	// init service

// 	nameNodeCache := make(map[int]*pkgEt.ServiceDiscovery)
// 	nameNodeCacheMtx := &sync.Mutex{}

// 	serviceRegistry := pkgSvc.NewServiceRegistry(cfg.ServiceRegistryConfig)
// 	dataNodeService := svcImpl.NewDataNodeService()
// 	nameNodeService := svcImpl.NewNameNodeService(&svcImpl.NameNodeServiceDto{
// 		ServiceRegistry: &serviceRegistry,
// 		NameNodeCache:   nameNodeCache,
// 		Mtx:             nameNodeCacheMtx,
// 	})

// 	// init usecase
// 	writeUC := ucImpl.NewWriteUsecase(&dataNodeService, &nameNodeService)
// }
