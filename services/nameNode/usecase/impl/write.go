package impl

import (
	"context"
	"fmt"
	"sync"

	"github.com/hdkef/hadoop/services/nameNode/config"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/repository"
	"github.com/hdkef/hadoop/services/nameNode/service"
	svcImpl "github.com/hdkef/hadoop/services/nameNode/service/impl"
	"github.com/hdkef/hadoop/services/nameNode/usecase"
)

type WriteRequestUsecaseImpl struct {
	metadataRepo    repository.MetadataRepo
	nodeStorageRepo repository.NodeStorageRepo
	jobQueueRepo    repository.JobQueueRepo
	iNodeRepo       repository.INodeRepo
	serviceRegistry service.ServiceRegistry
	dataNodeCache   map[string]*entity.ServiceDiscovery
	mtx             *sync.Mutex
	cfg             *config.Config
	nodeAllocator   service.NodeAllocator
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

func NewWriteUsecase() usecase.WriteRequestUsecase {
	return &WriteRequestUsecaseImpl{
		dataNodeCache: make(map[string]*entity.ServiceDiscovery),
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
func (w *WriteRequestUsecaseImpl) CreateRequest(ctx context.Context, dto *entity.CreateReqDto) (err error) {

	replTarget := w.cfg.ReplicationTarget
	blockSplitTarget := w.cfg.BlockSplitTarget
	if dto.BlockSplitTarget != 0 {
		blockSplitTarget = dto.BlockSplitTarget
	}
	if dto.ReplicationTarget != 0 {
		replTarget = dto.ReplicationTarget
	}

	// check metadata (cache / postgres)

	// check available dataNode (consul)
	var svd []*entity.ServiceDiscovery
	if len(w.dataNodeCache) == 0 {
		svd, err = w.serviceRegistry.GetAll(ctx, "dataNode", "")
		if err != nil {
			return err
		}
		w.mtx.Lock()
		defer w.mtx.Unlock()
		for _, v := range svd {
			w.dataNodeCache[v.GetID()] = v
		}
	}

	if len(svd) < int(replTarget) {
		return fmt.Errorf("available %d nodes, replication target %d", len(svd), replTarget)
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
	_, _, err = w.nodeAllocator.Allocate(nodeStorages, replTarget, blockSplitTarget, dto.FileSize)
	if err != nil {
		return err
	}

	// generate new job queue (dragonFly)

	// create transaction logs

	panic("unimplemented")
}
