package impl

import (
	"github.com/hdkef/hadoop/services/dataNode/config"
	"github.com/hdkef/hadoop/services/dataNode/service"
)

type ServiceRegistry struct {
}

func NewServiceRegistry() service.ServiceRegistry {
	return &ServiceRegistry{}
}

// RegisterNode implements service.ServiceRegistry.
func (s *ServiceRegistry) RegisterDataNode(cfg *config.Config) {
	panic("unimplemented")
}
