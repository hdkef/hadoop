package service

import (
	"context"

	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type DataNodeService interface {
	QueryStorage(ctx context.Context, svd *entity.ServiceDiscovery) (*entity.NodeStorage, error)
}
