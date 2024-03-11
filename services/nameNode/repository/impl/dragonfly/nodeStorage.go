package dragonfly

import (
	"context"

	pkgRepo "github.com/hdkef/hadoop/pkg/repository"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/repository"
)

type NodeStorageRepo struct {
	client pkgRepo.KeyValueRepository
}

// DecrLeaseStorage implements repository.NodeStorageRepo.
func (n *NodeStorageRepo) DecrLeaseStorage(ctx context.Context, et *entity.NodeStorage, amount int) error {
	panic("unimplemented")
}

// GetNodeStorage implements repository.NodeStorageRepo.
func (n *NodeStorageRepo) GetNodeStorage(ctx context.Context, nodeID string) (*entity.NodeStorage, error) {
	panic("unimplemented")
}

// IncrLeaseStorage implements repository.NodeStorageRepo.
func (n *NodeStorageRepo) IncrLeaseStorage(ctx context.Context, et *entity.NodeStorage, amount int) error {
	panic("unimplemented")
}

// SetNodeStorage implements repository.NodeStorageRepo.
func (n *NodeStorageRepo) SetNodeStorage(ctx context.Context, et *entity.NodeStorage) error {
	panic("unimplemented")
}

func NewNodeStorage(client pkgRepo.KeyValueRepository) repository.NodeStorageRepo {
	return &NodeStorageRepo{
		client: client,
	}
}
