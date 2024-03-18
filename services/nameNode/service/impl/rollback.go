package impl

import (
	"context"
	"sync"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	pkgSvc "github.com/hdkef/hadoop/pkg/services"

	"github.com/hdkef/hadoop/pkg/logger"
	pkgRepo "github.com/hdkef/hadoop/pkg/repository/transactionable"

	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/repository"
	"github.com/hdkef/hadoop/services/nameNode/service"
	"golang.org/x/sync/errgroup"
)

type RollbackServiceDto struct {
	DataNodeCache       map[string]*pkgEt.ServiceDiscovery
	Mtx                 *sync.Mutex
	TransactionsRepo    *repository.TransactionsRepo
	DataNodeService     *service.DataNodeService
	MetadataRepo        *repository.MetadataRepo
	TransactionInjector *pkgRepo.TransactionInjector
	ServiceRegistry     *pkgSvc.ServiceRegistry
	INodeRepo           *repository.INodeRepo
}

type RollbackService struct {
	dataNodeCache       map[string]*pkgEt.ServiceDiscovery
	mtx                 *sync.Mutex
	transactionsRepo    repository.TransactionsRepo
	dataNodeService     service.DataNodeService
	metadataRepo        repository.MetadataRepo
	transactionInjector *pkgRepo.TransactionInjector
	serviceRegistry     pkgSvc.ServiceRegistry
	iNodeRepo           repository.INodeRepo
}

// Rollback implements service.RollbackService.
func (r *RollbackService) Rollback(ctx context.Context, tx *entity.Transactions) error {

	// increment rollback retries
	err := r.transactionsRepo.IncrementRollbackRetries(ctx, tx.GetID(), nil)
	if err != nil {
		logger.LogError(err)
		return err
	}

	errGroup := &errgroup.Group{}

	// execute rollback

	// remove i_node_blocks
	r.iNodeRepo.Delete(ctx, tx.GetMetadata().GetINodeID(), nil)

	// remove metadata
	r.metadataRepo.Delete(ctx, tx.GetMetadata(), nil)

	r.mtx.Lock()
	if len(r.dataNodeCache) == 0 {
		// get dataNode via service discovery
		svds, err := r.serviceRegistry.GetAll(ctx, "dataNode", "")
		if err != nil {
			logger.LogError(err)
			return err
		}
		// delete old registry
		for key := range r.dataNodeCache {
			delete(r.dataNodeCache, key)
		}
		// set new registry
		for _, v := range svds {
			r.dataNodeCache[v.GetID()] = v
		}
	}
	r.mtx.Unlock()

	// remove files in dataNode
	blocks := tx.GetBlockTaret()
	for _, v := range blocks {
		for _, k := range v.NodeIDs {
			_, check := r.dataNodeCache[k]
			if !check {
				continue
			}

			dto := &entity.RollbackDto{}
			dto.SetBlockID(v.ID)
			dto.SetINodeID(tx.GetMetadata().GetINodeID())
			dto.SetNodeAddress(r.dataNodeCache[k].GetAddress())
			dto.SetNodePort(r.dataNodeCache[k].GetPort())
			dto.SetNodeID(k)

			errGroup.Go(func() error {
				return r.dataNodeService.Rollback(ctx, dto)
			})
		}
	}

	err = errGroup.Wait()
	if err != nil {
		logger.LogError(err)
		return err
	}

	// remove i_nodes_blocks

	// rolled back transaction
	err = r.transactionsRepo.RolledBack(ctx, tx.GetID(), nil)
	if err != nil {
		logger.LogError(err)
		return err
	}

	return nil
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

	if dto.TransactionInjector == nil {
		panic("transaction injector is nil")
	}

	if dto.INodeRepo == nil {
		panic("iNodeRepo is nil")
	}

	return &RollbackService{
		dataNodeCache:       dto.DataNodeCache,
		mtx:                 dto.Mtx,
		transactionsRepo:    *dto.TransactionsRepo,
		dataNodeService:     *dto.DataNodeService,
		metadataRepo:        *dto.MetadataRepo,
		transactionInjector: dto.TransactionInjector,
		serviceRegistry:     *dto.ServiceRegistry,
		iNodeRepo:           *dto.INodeRepo,
	}
}
