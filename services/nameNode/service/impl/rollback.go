package impl

import (
	"context"
	"sync"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/repository"
	"github.com/hdkef/hadoop/services/nameNode/service"
	"golang.org/x/sync/errgroup"
)

type RollbackServiceDto struct {
	DataNodeCache    map[string]*pkgEt.ServiceDiscovery
	Mtx              *sync.Mutex
	TransactionsRepo *repository.TransactionsRepo
	DataNodeService  *service.DataNodeService
	MetadataRepo     *repository.MetadataRepo
}

type RollbackService struct {
	dataNodeCache    map[string]*pkgEt.ServiceDiscovery
	mtx              *sync.Mutex
	transactionsRepo repository.TransactionsRepo
	dataNodeService  service.DataNodeService
	metadataRepo     repository.MetadataRepo
}

// Rollback implements service.RollbackService.
func (r *RollbackService) Rollback(ctx context.Context, tx *entity.Transactions) error {
	errGroup := &errgroup.Group{}

	// execute rollback

	// remove metadata
	r.metadataRepo.Delete(ctx, tx.GetMetadata(), nil)

	// remove files in dataNode
	blocks := tx.GetBlockTaret()
	for _, v := range blocks {
		for _, k := range v.NodeIDs {

			r.mtx.Lock()
			dto := &entity.RollbackDto{}
			dto.SetBlockID(v.ID)
			dto.SetINodeID(tx.GetMetadata().GetINodeID())
			dto.SetNodeAddress(r.dataNodeCache[k].GetAddress())
			dto.SetNodePort(r.dataNodeCache[k].GetPort())
			dto.SetNodeID(k)
			r.mtx.Unlock()

			errGroup.Go(func() error {
				return r.dataNodeService.Rollback(ctx, dto)
			})
		}
	}

	err := errGroup.Wait()
	if err != nil {
		return err
	}

	// rolled back transaction
	return r.transactionsRepo.RolledBack(ctx, tx.GetID(), nil)
}

func NewRollbackService(dto *RollbackServiceDto) service.RollbackService {

	if dto.DataNodeCache == nil {
		panic("dataNodeCache is nil")
	}

	if dto.Mtx == nil {
		panic("mtx is nil")
	}

	if dto.TransactionsRepo == nil {
		panic("transactionRepo is nil")
	}

	if dto.DataNodeService == nil {
		panic("dataNodeService is nil")
	}

	if dto.MetadataRepo == nil {
		panic("metadataRepo is nil")
	}

	return &RollbackService{
		dataNodeCache:    dto.DataNodeCache,
		mtx:              dto.Mtx,
		transactionsRepo: *dto.TransactionsRepo,
		dataNodeService:  *dto.DataNodeService,
		metadataRepo:     *dto.MetadataRepo,
	}
}
