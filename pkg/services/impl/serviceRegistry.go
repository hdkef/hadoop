package impl

import (
	"context"
	"fmt"
	"log"

	consul "github.com/hashicorp/consul/api"
	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	pkgSvc "github.com/hdkef/hadoop/pkg/services"
)

type ServiceRegistryConfig struct {
}

type ServiceRegistry struct {
	c *consul.Client
}

// GetAll implements service.ServiceRegistry.
func (s *ServiceRegistry) GetAll(ctx context.Context, servicesName string, tag string) ([]*pkgEt.ServiceDiscovery, error) {
	passingOnly := true
	addrs, _, err := s.c.Health().Service(servicesName, tag, passingOnly, nil)
	if len(addrs) == 0 && err == nil {
		return nil, fmt.Errorf("service ( %s ) was not found", servicesName)
	}
	if err != nil {
		return nil, err
	}

	svd := []*pkgEt.ServiceDiscovery{}

	for _, v := range addrs {

		newEntry := &pkgEt.ServiceDiscovery{}
		newEntry.SetAddress(v.Node.Address)
		newEntry.SetID(v.Node.ID)
		newEntry.SetPort(uint32(v.Service.Port))
		newEntry.SetServices(v.Service.Service)

		svd = append(svd, newEntry)
	}

	return svd, nil
}

func (s *ServiceRegistry) RegisterDataNode(id string, serviceName string, grpcport int, address string) {

	config := consul.DefaultConfig()
	cl, err := consul.NewClient(config)
	if err != nil {
		panic(err)
	}

	registeration := &consul.AgentServiceRegistration{
		ID:      id,
		Name:    "dataNode",
		Port:    grpcport,
		Address: address,
		Check: &consul.AgentServiceCheck{
			GRPC:     fmt.Sprintf("%s/%s", address, "Check"),
			Interval: "1s",
			Timeout:  "30s",
		},
	}

	regiErr := cl.Agent().ServiceRegister(registeration)
	if regiErr != nil {
		log.Panic(regiErr)
		log.Printf("Failed to register service: %s:%v ", address, grpcport)
	} else {
		log.Printf("successfully register service: %s:%v", address, grpcport)
	}

}

func NewServiceRegistryConfig() *ServiceRegistryConfig {
	return &ServiceRegistryConfig{}
}

func NewServiceRegistry(cfg *ServiceRegistryConfig) pkgSvc.ServiceRegistry {

	if cfg == nil {
		panic("config is nil")
	}

	config := consul.DefaultConfig()
	c, err := consul.NewClient(config)
	if err != nil {
		panic(err)
	}

	return &ServiceRegistry{
		c: c,
	}
}
