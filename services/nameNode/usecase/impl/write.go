package impl

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	"github.com/hdkef/hadoop/services/nameNode/config"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/repository"
	"github.com/hdkef/hadoop/services/nameNode/service"
	svcImpl "github.com/hdkef/hadoop/services/nameNode/service/impl"
	"github.com/hdkef/hadoop/services/nameNode/usecase"
)

type WriteRequestUsecaseImpl struct {
	metadataRepo     repository.MetadataRepo
	nodeStorageRepo  repository.NodeStorageRepo
	iNodeRepo        repository.INodeRepo
	serviceRegistry  service.ServiceRegistry
	dataNodeCache    map[string]*entity.ServiceDiscovery
	transactionsRepo repository.TransactionsRepo
	mtx              *sync.Mutex
	cfg              *config.Config
	nodeAllocator    service.NodeAllocator
}

// CheckDataNode implements usecase.WriteRequestUsecase.
func (w *WriteRequestUsecaseImpl) CheckDataNode(ctx context.Context) error {
	svd, err := w.serviceRegistry.GetAll(ctx, "dataNode", "")
	if err != nil {
		return err
	}
	w.mtx.Lock()
	defer w.mtx.Unlock()
	for _, v := range svd {
		w.dataNodeCache[v.GetID()] = v
	}
	return nil
}

func NewWriteUsecase(dataNodeCache map[string]*entity.ServiceDiscovery) usecase.WriteRequestUsecase {
	return &WriteRequestUsecaseImpl{
		dataNodeCache: dataNodeCache,
		mtx:           &sync.Mutex{},
		nodeAllocator: svcImpl.NewNodeAllocator(),
	}
}

// CommitCreateRequest implements usecase.WriteRequestUsecase.
func (w *WriteRequestUsecaseImpl) CommitCreateRequest(ctx context.Context) error {

	// create iNodeblockids

	// update transaction checkpoint

	panic("unimplemented")
}

// WriteRequest implements usecase.WriteRequestUsecase.
func (w *WriteRequestUsecaseImpl) CreateRequest(ctx context.Context, dto *entity.CreateReqDto) (res []*pkgEt.QueryNodeTarget, err error) {

	replTarget := w.cfg.ReplicationTarget
	blockSplitTarget := w.cfg.BlockSplitTarget
	leaseTimeInSec := uint64(w.cfg.MinLeaseTime.Seconds())
	if dto.BlockSplitTarget != 0 {
		blockSplitTarget = dto.BlockSplitTarget
	}
	if dto.ReplicationTarget != 0 {
		replTarget = dto.ReplicationTarget
	}
	if dto.LeaseTimeInSec != 0 {
		leaseTimeInSec = dto.LeaseTimeInSec
	}

	// TODO:
	// check metadata (cache / postgres)
	metadata := &entity.Metadata{}

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
		} else {
			nodeStorages = append(nodeStorages, nd)
		}
	}

	// allocate targetNode per block
	var blockTargets []*entity.BlockTarget
	blockTargets, nodeStorages, err = w.nodeAllocator.Allocate(nodeStorages, replTarget, blockSplitTarget, dto.FileSize)
	if err != nil {
		return nil, err
	}

	// create transaction logs

	transactions := &entity.Transactions{}
	newID := uuid.New()
	transactions.SetID(newID)
	transactions.SetMetadata(metadata)
	transactions.SetAction(entity.TRANSACTION_ACTION_CREATE)
	transactions.SetLeaseTimeInSecond(leaseTimeInSec)

	transactionNodeInfo := []*entity.TransactionNodeInfo{}

	allBlockIDs := []uuid.UUID{}

	for _, v := range blockTargets {
		allBlockIDs = append(allBlockIDs, v.ID)
	}

	for _, v := range blockTargets {

		replNodeTarget := []*pkgEt.NodeTarget{}

		for _, k := range v.NodeIDs {
			// domain transaction
			newTr := &entity.TransactionNodeInfo{}
			newTr.SetAddress(w.dataNodeCache[k].GetAddress())
			newTr.SetGRPCPort(w.dataNodeCache[k].GetPort())
			newTr.SetBlockID(v.ID)
			newTr.SetNodeID(k)
			newTr.SetFileSize(v.Size)
			transactionNodeInfo = append(transactionNodeInfo, newTr)

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

	transactions.SetTransactionNodeInfo(transactionNodeInfo)

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
