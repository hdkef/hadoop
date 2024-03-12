package impl

import (
	"context"
	"fmt"

	consul "github.com/hashicorp/consul/api"
	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	pkgSvc "github.com/hdkef/hadoop/pkg/services"
)

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

func NewServiceRegistry() pkgSvc.ServiceRegistry {
	config := consul.DefaultConfig()
	c, err := consul.NewClient(config)
	if err != nil {
		panic(err)
	}

	return &ServiceRegistry{
		c: c,
	}
}
