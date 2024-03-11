package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	"github.com/hdkef/hadoop/services/nameNode/entity"
)

// WriteRequest implements usecase.WriteRequestUsecase.
func (w *WriteRequestUsecaseImpl) CreateRequest(ctx context.Context, dto *entity.CreateReqDto) (res []*pkgEt.QueryNodeTarget, err error) {

	replTarget := w.cfg.ReplicationTarget
	blockSplitTarget := w.cfg.BlockSplitTarget
	leaseTimeInSec := uint64(w.cfg.MinLeaseTime.Seconds())
	if dto.GetBlockSplitTarget() != 0 {
		blockSplitTarget = dto.GetBlockSplitTarget()
	}
	if dto.GetReplicationTarget() != 0 {
		replTarget = dto.GetReplicationTarget()
	}
	if dto.GetLeaseTimeInSec() != 0 {
		leaseTimeInSec = dto.GetLeaseTimeInSec()
	}

	// check parentPath
	exist := w.metadataRepo.CheckPath(ctx, dto.GetParentPath())
	if !exist {
		return nil, errors.New("metadata parent path is not exist")
	}

	// check path
	exist = w.metadataRepo.CheckPath(ctx, dto.GetPath())
	if exist {
		return nil, errors.New("metadata in that path is already exist")
	}
	metadata := &entity.Metadata{}
	metadata.SetPath(dto.GetPath())
	metadata.SetParentPath(dto.GetParentPath())
	metadata.SetType(entity.METADATA_TYPE_FILE)
	metadata.SetINodeID(uuid.New())

	// check available dataNode (consul)
	var svd []*entity.ServiceDiscovery
	if len(w.dataNodeCache) == 0 {
		svd, err = w.serviceRegistry.GetAll(ctx, "dataNode", "")
		if err != nil {
			return nil, err
		}
		w.mtx.Lock()
		defer w.mtx.Unlock()
		for _, v := range svd {
			w.dataNodeCache[v.GetID()] = v
		}
	}

	if len(svd) < int(replTarget) {
		return nil, fmt.Errorf("available %d nodes, replication target %d", len(svd), replTarget)
	}

	// check available space dataNode (query)
	nodeStorages := make([]*entity.NodeStorage, 0)

	for _, v := range svd {

		nd, err := w.nodeStorageRepo.GetNodeStorage(ctx, v.GetID())
		if err != nil || nd == nil {
			// if in cache not exist, try query to dataNode
			qNd, err := w.dataNodeService.QueryStorage(ctx, v)
			if err != nil {
				continue
			}
			nodeStorages = append(nodeStorages, qNd)
		} else {
			nodeStorages = append(nodeStorages, nd)
		}
	}

	// allocate targetNode per block
	var blockTargets []*entity.BlockTarget
	blockTargets, nodeStorages, err = w.nodeAllocator.Allocate(nodeStorages, replTarget, blockSplitTarget, dto.GetFileSize())
	if err != nil {
		return nil, err
	}

	// create transaction logs

	transactions := &entity.Transactions{}
	transactions.SetID(uuid.New())
	transactions.SetMetadata(metadata)
	transactions.SetAction(entity.TRANSACTION_ACTION_CREATE)
	transactions.SetLeaseTimeInSecond(leaseTimeInSec)

	allBlockIDs := []uuid.UUID{}

	for _, v := range blockTargets {
		allBlockIDs = append(allBlockIDs, v.ID)
	}

	for _, v := range blockTargets {

		replNodeTarget := []*pkgEt.NodeTarget{}

		for _, k := range v.NodeIDs {
			newNodeTarget := &pkgEt.NodeTarget{}
			newNodeTarget.SetBlockID(v.ID)
			newNodeTarget.SetNodeAddress(w.dataNodeCache[k].GetAddress())
			newNodeTarget.SetNodeGrpcPort(w.dataNodeCache[k].GetPort())
			newNodeTarget.SetNodeID(k)
			replNodeTarget = append(replNodeTarget, newNodeTarget)
		}

		// domain query
		q := &pkgEt.QueryNodeTarget{}
		q.SetAllBlockID(allBlockIDs)
		q.SetINodeID(metadata.GetINodeID())
		q.SetTransactionID(transactions.GetID())
		q.SetNodeTargets(replNodeTarget)
		res = append(res, q)
	}

	transactions.SetBlockTarget(blockTargets)

	err = w.transactionsRepo.Add(ctx, transactions)
	if err != nil {
		return nil, err
	}

	// update leaseStorage in nodeStorage repo
	for _, v := range nodeStorages {
		if v.GetLeasedUsedStorageChanged() {
			err = w.nodeStorageRepo.SetNodeStorage(ctx, v)
			if err != nil {
				return nil, err
			}
		}
	}

	// respond

	return
}
