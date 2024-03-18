package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	"github.com/hdkef/hadoop/pkg/logger"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"golang.org/x/sync/errgroup"
)

// WriteRequest implements usecase.WriteRequestUsecase.
func (w *WriteRequestUsecaseImpl) CreateRequest(ctx context.Context, dto *pkgEt.CreateReqDto) (*pkgEt.QueryNodeTarget, error) {

	replTarget := w.cfg.ReplicationTarget
	blockSplitTarget := w.cfg.BlockSplitTarget
	leaseTimeInSec := uint32(w.cfg.MinLeaseTime.Seconds())
	if dto.GetBlockSplitTarget() != 0 {
		blockSplitTarget = dto.GetBlockSplitTarget()
	}
	if dto.GetReplicationTarget() != 0 {
		replTarget = dto.GetReplicationTarget()
	}
	if dto.GetLeaseTimeInSec() != 0 {
		leaseTimeInSec = dto.GetLeaseTimeInSec()
	}

	errGroup, c := errgroup.WithContext(ctx)

	// check parentPath
	errGroup.Go(func() error {
		exist := w.metadataRepo.CheckPath(c, dto.GetParentPath(), nil)
		if !exist {
			return errors.New("metadata parent path is not exist")
		}
		return nil
	})

	// check path
	errGroup.Go(func() error {
		exist := w.metadataRepo.CheckPath(c, dto.GetPath(), nil)
		if exist {
			return errors.New("metadata in that path is already exist")
		}
		return nil
	})

	err := errGroup.Wait()
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	metadata := &entity.Metadata{}
	metadata.SetPath(dto.GetPath())
	metadata.SetParentPath(dto.GetParentPath())
	metadata.SetType(entity.METADATA_TYPE_FILE)
	metadata.SetINodeID(uuid.New())
	metadata.SetHash(dto.GetHash())

	// check available dataNode (consul)
	var svd []*pkgEt.ServiceDiscovery
	w.mtx.Lock()
	if len(w.dataNodeCache) == 0 {
		svd, err = w.serviceRegistry.GetAll(ctx, "dataNode", "")
		if err != nil {
			logger.LogError(err)
			return nil, err
		}

		for _, v := range svd {
			w.dataNodeCache[v.GetID()] = v
		}
	}
	w.mtx.Unlock()

	if len(w.dataNodeCache) < int(replTarget) {
		return nil, fmt.Errorf("available %d nodes, replication target %d", len(svd), replTarget)
	}

	// check available space dataNode (query)
	nodeStorages := make([]*entity.NodeStorage, 0)

	for _, v := range w.dataNodeCache {

		nd := &entity.NodeStorage{}
		nd.SetNodeID(v.GetID())

		err := w.nodeStorageRepo.GetNodeStorage(ctx, nd)
		if err != nil {
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
		logger.LogError(err)
		return nil, err
	}

	// create transaction logs

	transactions := &entity.Transactions{}
	transactions.SetID(uuid.New())
	transactions.SetAction(entity.TRANSACTION_ACTION_CREATE)
	transactions.SetLeaseTimeInSecond(leaseTimeInSec)

	allBlockIDs := []uuid.UUID{}

	for _, v := range blockTargets {
		allBlockIDs = append(allBlockIDs, v.ID)
	}

	w.mtx.Lock()
	replNodeTarget := []*pkgEt.NodeTarget{}
	for _, v := range blockTargets {

		for _, k := range v.NodeIDs {
			newNodeTarget := &pkgEt.NodeTarget{}
			newNodeTarget.SetBlockID(v.ID)
			newNodeTarget.SetNodeAddress(w.dataNodeCache[k].GetAddress())
			newNodeTarget.SetNodeGrpcPort(w.dataNodeCache[k].GetPort())
			newNodeTarget.SetNodeID(k)
			replNodeTarget = append(replNodeTarget, newNodeTarget)
		}
	}
	w.mtx.Unlock()

	transactions.SetBlockTarget(blockTargets)
	metadata.SetAllBlockIds(allBlockIDs)
	transactions.SetMetadata(metadata)

	trId, err := w.transactionsRepo.Add(ctx, transactions, nil)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}
	transactions.SetID(trId)

	// update leaseStorage in nodeStorage repo
	for _, v := range nodeStorages {
		if v.GetLeasedUsedStorageChanged() {
			err = w.nodeStorageRepo.SetNodeStorage(ctx, v)
			if err != nil {
				logger.LogError(err)
				return nil, err
			}
		}
	}

	// respond
	// domain query
	res := &pkgEt.QueryNodeTarget{}

	res.SetAllBlockID(allBlockIDs)
	res.SetINodeID(metadata.GetINodeID())
	res.SetTransactionID(transactions.GetID())
	res.SetNodeTargets(replNodeTarget)
	res.SetReplicationFactor(replTarget)

	return res, nil
}
