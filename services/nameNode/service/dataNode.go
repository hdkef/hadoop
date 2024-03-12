package service

import (
	"context"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type DataNodeService interface {
	QueryStorage(ctx context.Context, svd *pkgEt.ServiceDiscovery) (*entity.NodeStorage, error)
	Rollback(ctx context.Context, dto *entity.RollbackDto) error
}
