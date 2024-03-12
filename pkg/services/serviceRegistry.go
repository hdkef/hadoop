package service

import (
	"context"

	"github.com/hdkef/hadoop/pkg/entity"
)

type ServiceRegistry interface {
	GetAll(ctx context.Context, servicesName string, tag string) ([]*entity.ServiceDiscovery, error)
}
