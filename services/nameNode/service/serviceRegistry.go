package service

import (
	"context"

	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type ServiceRegistry interface {
	GetAll(ctx context.Context, servicesName string, tag string) ([]*entity.ServiceDiscovery, error)
}
