package main

import (
	"fmt"
	"log"
	"net"

	pkgSvc "github.com/hdkef/hadoop/pkg/services/impl"
	"github.com/hdkef/hadoop/services/dataNode/config"
	"github.com/hdkef/hadoop/services/dataNode/delivery/grpc"
	serviceImpl "github.com/hdkef/hadoop/services/dataNode/service/impl"
	usecaseImpl "github.com/hdkef/hadoop/services/dataNode/usecase/impl"
)

func main() {

	// init config
	cfg := config.NewConfig()

	// spin up tcp listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GrpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// register dataNode to service registry
	svcRegistry := pkgSvc.NewServiceRegistry(cfg.ServiceRegistryConfig)
	svcRegistry.RegisterDataNode(cfg.NodeId, "dataNode", cfg.GrpcPort, cfg.Address)

	// init usecase
	dataNodeService := serviceImpl.NewDataNodeService()
	writeUC := usecaseImpl.NewWriteUsecase(cfg, &dataNodeService)

	log.Printf("dataNode %s will run on address %s & grpc port %d", cfg.NodeId, cfg.Address, cfg.GrpcPort)

	// spin up grpc server
	sv := grpc.NewGrpcHandler(cfg, writeUC)
	if err := sv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
