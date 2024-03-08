package impl

import (
	"fmt"
	"log"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/hdkef/hadoop/services/dataNode/config"
	"github.com/hdkef/hadoop/services/dataNode/service"
)

type ServiceRegistry struct {
	cfg *config.Config
}

func NewServiceRegistry(cfg *config.Config) service.ServiceRegistry {
	return &ServiceRegistry{
		cfg: cfg,
	}
}

// RegisterNode implements service.ServiceRegistry.
func (s *ServiceRegistry) RegisterDataNode() {

	cfg := s.cfg

	if cfg == nil {
		panic("nil config")
	}

	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		panic(err)
	}

	registeration := &consulapi.AgentServiceRegistration{
		ID:      cfg.NodeId,
		Name:    "dataNode",
		Port:    cfg.GrpcPort,
		Address: cfg.Address,
		Check: &consulapi.AgentServiceCheck{
			GRPC:     fmt.Sprintf("%s/%s", cfg.Address, "Check"),
			Interval: "1s",
			Timeout:  "30s",
		},
	}

	regiErr := consul.Agent().ServiceRegister(registeration)
	if regiErr != nil {
		log.Panic(regiErr)
		log.Printf("Failed to register service: %s:%v ", cfg.Address, cfg.GrpcPort)
	} else {
		log.Printf("successfully register service: %s:%v", cfg.Address, cfg.GrpcPort)
	}

}
