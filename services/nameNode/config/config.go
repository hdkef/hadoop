package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	pkgRepoDragonfly "github.com/hdkef/hadoop/pkg/repository/dragonfly"
	pkgRepoPostgres "github.com/hdkef/hadoop/pkg/repository/postgres"
	pkgSvc "github.com/hdkef/hadoop/pkg/services/impl"
)

const (
	REPLICATION_TARGET = "REPLICATION_TARGET"
	BLOCK_SPLIT_TARGET = "BLOCK_SPLIT_TARGET"
	MIN_LEASE_TIME     = "MIN_LEASE_TIME"
	NAME_NODE_PORT     = "NAME_NODE_PORT"
	NAME_NODE_ADDRESS  = "NAME_NODE_ADDRESS"
	NODE_ID            = "NODE_ID"
)

type Config struct {
	ReplicationTarget     uint32
	BlockSplitTarget      uint32
	MinLeaseTime          time.Duration
	DragonFlyConfig       *pkgRepoDragonfly.DragonFlyConfig
	PostgresConfig        *pkgRepoPostgres.PostgresConfig
	NameNodePort          uint32
	NameNodeAddress       string
	ServiceRegistryConfig *pkgSvc.ServiceRegistryConfig
	NodeID                string
}

func NewConfig() *Config {

	replTarget := os.Getenv(REPLICATION_TARGET)

	if replTarget == "" {
		panic(fmt.Sprintf("%s env not found", REPLICATION_TARGET))
	}

	replTargetVal, err := strconv.Atoi(replTarget)
	if err != nil {
		panic(fmt.Sprintf("%s %s", REPLICATION_TARGET, err.Error()))
	}

	blockSplitTarget := os.Getenv(BLOCK_SPLIT_TARGET)

	if blockSplitTarget == "" {
		panic(fmt.Sprintf("%s env not found", BLOCK_SPLIT_TARGET))
	}

	blockSplitTargetVal, err := strconv.Atoi(blockSplitTarget)
	if err != nil {
		panic(fmt.Sprintf("%s %s", BLOCK_SPLIT_TARGET, err.Error()))
	}

	minLeaseTime := os.Getenv(MIN_LEASE_TIME)

	if minLeaseTime == "" {
		panic(fmt.Sprintf("%s env not found", MIN_LEASE_TIME))
	}

	minLeaseTimeVal, err := time.ParseDuration(minLeaseTime)
	if err != nil {
		panic(fmt.Sprintf("%s %s", MIN_LEASE_TIME, err.Error()))
	}

	nameNodePort := os.Getenv(NAME_NODE_PORT)

	if nameNodePort == "" {
		panic(fmt.Sprintf("%s env not found", NAME_NODE_PORT))
	}

	nameNodePortVal, err := strconv.Atoi(nameNodePort)
	if err != nil {
		panic(fmt.Sprintf("%s %s", NAME_NODE_PORT, err.Error()))
	}

	address := os.Getenv(NAME_NODE_ADDRESS)
	if address == "" {
		panic(fmt.Sprintf("%s env not found", NAME_NODE_ADDRESS))
	}

	nodeId := os.Getenv(NODE_ID)
	if nodeId == "" {
		panic(fmt.Sprintf("%s env not found", NODE_ID))
	}

	return &Config{
		ReplicationTarget:     uint32(replTargetVal),
		BlockSplitTarget:      uint32(blockSplitTargetVal),
		MinLeaseTime:          minLeaseTimeVal,
		PostgresConfig:        pkgRepoPostgres.NewPostgresConfig(),
		DragonFlyConfig:       pkgRepoDragonfly.NewDragonFlyConfig(),
		NameNodePort:          uint32(nameNodePortVal),
		ServiceRegistryConfig: pkgSvc.NewServiceRegistryConfig(),
		NameNodeAddress:       address,
		NodeID:                nodeId,
	}
}
