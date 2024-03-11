package repository

import (
	"context"

	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type NodeStorageRepo interface {
	IncrLeaseStorage(ctx context.Context, et *entity.NodeStorage, amount int) error
	DecrLeaseStorage(ctx context.Context, et *entity.NodeStorage, amount int) error
	GetNodeStorage(ctx context.Context, nodeID string) (*entity.NodeStorage, error)
	SetNodeStorage(ctx context.Context, et *entity.NodeStorage) error
}
