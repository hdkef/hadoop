package config

import (
	"fmt"
	"os"
	"strconv"

	pkgSvc "github.com/hdkef/hadoop/pkg/services/impl"
)

const (
	NODE_ID           = "NODE_ID"
	GRPC_PORT         = "GRPC_PORT"
	ADDRESS           = "ADDRESS"
	NAME_NODE_ADDRESS = "NAME_NODE_ADDRESS"
	NAME_NODE_PORT    = "NAME_NODE_PORT"
	STORAGE_ROOT      = "STORAGE_ROOT"
)

type Config struct {
	NodeId                string
	GrpcPort              int
	Address               string
	NameNodeAddress       string
	NameNodePort          int
	StorageRoot           string
	ServiceRegistryConfig *pkgSvc.ServiceRegistryConfig
}

func NewConfig() *Config {

	nodeId := os.Getenv(NODE_ID)
	if nodeId == "" {
		panic(fmt.Sprintf("%s env not found", NODE_ID))
	}

	grpcPort := os.Getenv(GRPC_PORT)
	if grpcPort == "" {
		panic(fmt.Sprintf("%s env not found", GRPC_PORT))
	}

	grpcPortVal, err := strconv.Atoi(grpcPort)
	if err != nil {
		panic(fmt.Sprintf("%s %s", GRPC_PORT, err.Error()))
	}

	address := os.Getenv(ADDRESS)
	if address == "" {
		panic(fmt.Sprintf("%s env not found", ADDRESS))
	}

	nameNodeAddress := os.Getenv(NAME_NODE_ADDRESS)
	if nameNodeAddress == "" {
		panic(fmt.Sprintf("%s env not found", NAME_NODE_ADDRESS))
	}

	nameNodePort := os.Getenv(NAME_NODE_PORT)
	if nameNodePort == "" {
		panic(fmt.Sprintf("%s env not found", NAME_NODE_PORT))
	}

	nameNodePortVal, err := strconv.Atoi(nameNodePort)
	if err != nil {
		panic(fmt.Sprintf("%s %s", NAME_NODE_PORT, err.Error()))
	}

	storageRoot := os.Getenv(STORAGE_ROOT)
	if storageRoot == "" {
		storageRoot = "/app"
	}

	return &Config{
		NodeId:                nodeId,
		GrpcPort:              grpcPortVal,
		Address:               address,
		NameNodeAddress:       nameNodeAddress,
		NameNodePort:          nameNodePortVal,
		StorageRoot:           storageRoot,
		ServiceRegistryConfig: pkgSvc.NewServiceRegistryConfig(),
	}
}
