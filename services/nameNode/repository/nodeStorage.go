package repository

import (
	"context"

	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type NodeStorageRepo interface {
	GetNodeStorage(ctx context.Context, et *entity.NodeStorage) error
	SetNodeStorage(ctx context.Context, et *entity.NodeStorage) error
}
