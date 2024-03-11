package service

import "github.com/hdkef/hadoop/services/nameNode/entity"

type NodeAllocator interface {
	Allocate(nodeStorage []*entity.NodeStorage, replicationTarget uint32, blockSplitTarget uint32, fileSize uint64) ([]*entity.BlockTarget, []*entity.NodeStorage, error)
}
