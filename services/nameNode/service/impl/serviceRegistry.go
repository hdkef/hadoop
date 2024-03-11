package impl

import (
	"context"
	"fmt"

	consul "github.com/hashicorp/consul/api"
	"github.com/hdkef/hadoop/services/nameNode/config"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/service"
)

type ServiceRegistry struct {
	cfg *config.Config
	c   *consul.Client
}

// GetAll implements service.ServiceRegistry.
func (s *ServiceRegistry) GetAll(ctx context.Context, servicesName string, tag string) ([]*entity.ServiceDiscovery, error) {
	passingOnly := true
	addrs, _, err := s.c.Health().Service(servicesName, tag, passingOnly, nil)
	if len(addrs) == 0 && err == nil {
		return nil, fmt.Errorf("service ( %s ) was not found", servicesName)
	}
	if err != nil {
		return nil, err
	}

	svd := []*entity.ServiceDiscovery{}

	for _, v := range addrs {

		newEntry := &entity.ServiceDiscovery{}
		newEntry.SetAddress(v.Node.Address)
		newEntry.SetID(v.Node.ID)
		newEntry.SetPort(uint32(v.Service.Port))
		newEntry.SetServices(v.Service.Service)

		svd = append(svd, newEntry)
	}

	return svd, nil
}

func NewServiceRegistry(cfg *config.Config) service.ServiceRegistry {
	config := consul.DefaultConfig()
	c, err := consul.NewClient(config)
	if err != nil {
		panic(err)
	}

	return &ServiceRegistry{
		cfg: cfg,
		c:   c,
	}
}
