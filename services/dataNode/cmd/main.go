package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hdkef/hadoop/services/dataNode/config"
	"github.com/hdkef/hadoop/services/dataNode/delivery/grpc"
	repoImpl "github.com/hdkef/hadoop/services/dataNode/repository/impl"
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
	svcRegistry := serviceImpl.NewServiceRegistry(cfg)
	svcRegistry.RegisterDataNode()

	// init repository
	kvRepo := repoImpl.NewkeyValueRepo(cfg.StorageRoot)
	defer kvRepo.CloseConn()

	// init usecase
	writeUC := usecaseImpl.NewWriteUsecase(cfg, kvRepo)

	log.Printf("dataNode %s will run on address %s & grpc port %d", cfg.NodeId, cfg.Address, cfg.GrpcPort)

	// spin up grpc server
	sv := grpc.NewGrpcHandler(cfg, writeUC)
	if err := sv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
