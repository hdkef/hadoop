package impl

import (
	"errors"
	"fmt"
	"sort"

	"github.com/google/uuid"
	"github.com/hdkef/hadoop/services/nameNode/config"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/service"
)

type ByTotalStorage []*entity.NodeStorage

func (a ByTotalStorage) Len() int      { return len(a) }
func (a ByTotalStorage) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTotalStorage) Less(i, j int) bool {

	// first prioritize lower allocated amount

	// second prioritize greater available spaces

	if a[i].GetAllocated() == a[j].GetAllocated() {
		return a[i].GetFreeSpace() > a[j].GetFreeSpace()
	}

	return a[i].GetAllocated() < a[j].GetAllocated()
}

type NodeAllocator struct {
	cfg *config.Config
}

// Allocate implements service.NodeAllocator.
func (n *NodeAllocator) Allocate(nodeStorage []*entity.NodeStorage, replicationTarget uint32, blockSplitTarget uint32, fileSize uint64) ([]*entity.BlockTarget, []*entity.NodeStorage, error) {

	if len(nodeStorage) < int(replicationTarget) {
		return nil, nil, fmt.Errorf("available %d nodes, replication target %d", len(nodeStorage), replicationTarget)
	}

	blocks := make([]*entity.BlockTarget, blockSplitTarget)

	// Allocate size for each block and assign blockID
	singleBlockSize := fileSize / uint64(blockSplitTarget)
	lastBlockSize := fileSize - singleBlockSize*(uint64(blockSplitTarget)-1)

	for i := 0; i < int(blockSplitTarget); i++ {

		size := singleBlockSize
		if i == int(blockSplitTarget)-1 {
			size = lastBlockSize
		}

		blocks[i] = &entity.BlockTarget{
			ID:   uuid.New(),
			Size: size,
		}

		// Allocate block to selected nodes
		// sort nodes
		sort.Sort(ByTotalStorage(nodeStorage))

		allocatedNodeId := make(map[string]bool)

		for i2 := 0; i2 < int(replicationTarget); i2++ {

			// take most prioritize and not yet allocated
			taken := nodeStorage[i2]

			if taken.GetFreeSpace() >= size && !allocatedNodeId[taken.GetNodeID()] {
				nodeStorage[i2].SetAllocated()
				// assign node
				nodeStorage[i2].IncrementLeaseStorage(size)
				nodeStorage[i2].SetLeasedUsedStorageChanged()
				blocks[i].NodeIDs = append(blocks[i].NodeIDs, taken.GetNodeID())
				allocatedNodeId[taken.GetNodeID()] = true
			} else {
				// if space not available, check next nodes
				nextIdx := i2 + 1

				for nextIdx != i2 {
					if nextIdx >= len(nodeStorage) {
						nextIdx = 0
					}

					taken = nodeStorage[nextIdx]
					if taken.GetFreeSpace() >= size && !allocatedNodeId[taken.GetNodeID()] {
						nodeStorage[nextIdx].SetAllocated()
						// assign node
						nodeStorage[nextIdx].IncrementLeaseStorage(size)
						nodeStorage[nextIdx].SetLeasedUsedStorageChanged()
						blocks[i].NodeIDs = append(blocks[i].NodeIDs, taken.GetNodeID())
						allocatedNodeId[taken.GetNodeID()] = true
						break
					} else {
						nextIdx++
					}
				}

				if nextIdx == i2 {
					return nil, nil, errors.New("not enough available nodes to allocate the block")
				}
			}
		}
	}

	return blocks, nodeStorage, nil
}

func NewNodeAllocator(cfg *config.Config) service.NodeAllocator {
	if cfg == nil {
		panic("nil config")
	}

	return &NodeAllocator{
		cfg: cfg,
	}
}
