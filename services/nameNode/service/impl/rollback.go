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

type RollbackService struct {
	dataNodeCache    map[string]*pkgEt.ServiceDiscovery
	mtx              *sync.Mutex
	transactionsRepo repository.TransactionsRepo
	dataNodeService  DataNodeService
	metadataRepo     repository.MetadataRepo
}

// Rollback implements service.RollbackService.
func (r *RollbackService) Rollback(ctx context.Context, tx *entity.Transactions) error {
	errGroup := &errgroup.Group{}

	// execute rollback

	// remove metadata
	errGroup.Go(func() error {
		return r.metadataRepo.Delete(ctx, tx.GetMetadata())
	})

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
	return r.transactionsRepo.RolledBack(ctx, tx.GetID())
}

func NewRollbackService() service.RollbackService {
	return &RollbackService{}
}
