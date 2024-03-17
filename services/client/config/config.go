package config

import (
	"fmt"
	"os"
	"strconv"

	pkgSvc "github.com/hdkef/hadoop/pkg/services/impl"
)

const (
	CLIENT_PORT = "CLIENT_PORT"
)

type Config struct {
	ClientPort            uint32
	ServiceRegistryConfig *pkgSvc.ServiceRegistryConfig
}

func NewConfigServer() *Config {

	clientPort := os.Getenv(CLIENT_PORT)
	if clientPort == "" {
		panic(fmt.Sprintf("%s env not found", CLIENT_PORT))
	}

	clientPortVal, err := strconv.Atoi(clientPort)
	if err != nil {
		panic(fmt.Sprintf("%s %s", CLIENT_PORT, err.Error()))
	}

	return &Config{
		ServiceRegistryConfig: pkgSvc.NewServiceRegistryConfig(),
		ClientPort:            uint32(clientPortVal),
	}
}
