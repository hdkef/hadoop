package dragonfly

import (
	"context"
	"errors"

	pkgRepo "github.com/hdkef/hadoop/pkg/repository"
	messageProto "github.com/hdkef/hadoop/proto/message"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/repository"
	"google.golang.org/protobuf/proto"
)

type NodeStorageRepo struct {
	client pkgRepo.KeyValueRepository
}

// GetNodeStorage implements repository.NodeStorageRepo.
func (n *NodeStorageRepo) GetNodeStorage(ctx context.Context, et *entity.NodeStorage) error {

	key := et.GenerateKey()

	data, err := n.client.Get(ctx, key)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return errors.New("data not found")
	}

	pbNodeStorage := &messageProto.NodeStorage{}

	err = proto.Unmarshal(data, pbNodeStorage)
	if err != nil {
		return err
	}

	et.SetActualUsedStorage(pbNodeStorage.ActualUsedStorage)
	et.SetNodeID(pbNodeStorage.NodeID)
	et.SetLeaseUsedStorage(pbNodeStorage.LeaseUsedStorage)
	et.SetTotalStorage(pbNodeStorage.TotalStorage)

	return nil
}

// SetNodeStorage implements repository.NodeStorageRepo.
func (n *NodeStorageRepo) SetNodeStorage(ctx context.Context, et *entity.NodeStorage) error {

	key := et.GenerateKey()

	pbNodeStorage := &messageProto.NodeStorage{
		NodeID:            et.GetNodeID(),
		LeaseUsedStorage:  et.GetLeaseUsedStorage(),
		ActualUsedStorage: et.GetActualUsedStorage(),
		TotalStorage:      et.GetTotalStorage(),
	}

	data, err := proto.Marshal(pbNodeStorage)
	if err != nil {
		return err
	}
	return n.client.Set(ctx, key, data, nil)
}

func NewNodeStorage(client pkgRepo.KeyValueRepository) repository.NodeStorageRepo {

	return &NodeStorageRepo{
		client: client,
	}
}
