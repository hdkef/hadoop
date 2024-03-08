package impl

import (
	"context"
	"errors"

	"github.com/hdkef/hadoop/services/dataNode/config"
	"github.com/hdkef/hadoop/services/dataNode/entity"
	"github.com/hdkef/hadoop/services/dataNode/service"
	"github.com/hdkef/hadoop/services/dataNode/usecase"

	pkgRepo "github.com/hdkef/hadoop/pkg/repository"
)

type WriteUsecaseImpl struct {
	cfg             *config.Config
	kvRepo          pkgRepo.KeyValueRepository
	dataNodeService service.DataNodeService
	nameNodeService service.NameNodeService
}

func NewWriteUsecase(cfg *config.Config, kvRepo pkgRepo.KeyValueRepository, dataNodeService service.DataNodeService, nameNodeService service.NameNodeService) usecase.WriteUsecase {

	if cfg == nil {
		panic("nil config")
	}

	return &WriteUsecaseImpl{
		cfg:             cfg,
		kvRepo:          kvRepo,
		dataNodeService: dataNodeService,
	}
}

// Write implements usecase.WriteUsecase.
func (w *WriteUsecaseImpl) Write(ctx context.Context, dto *entity.WriteDto) error {

	iNodeBlockId := entity.INodeBlockID{}
	iNodeBlockId.SetINodeID(dto.GetInodeID())
	iNodeBlockId.SetBlockID(dto.GetBlockID())

	// write file to storage
	err := iNodeBlockId.Write(w.cfg.StorageRoot, dto.GetBlocksData())
	if err != nil {
		return err
	}

	// create new k,v store inode_blockid
	err = w.kvRepo.Set(ctx, iNodeBlockId.GetKey(), iNodeBlockId.ToJSON(), nil)
	if err != nil {
		return err
	}

	// update jobQueue nameNode
	err = w.nameNodeService.UpdateJobQueue(ctx, dto, true)
	if err != nil {
		return err
	}

	// increment currentReplicaSet
	dto.IncrementCurrentReplicated()

	// set node status
	for i, v := range dto.GetReplicationNodeTarget() {
		if v.GetNodeID() == w.cfg.NodeId {
			v.SetReplicationStatusSuccess()
			dto.UpdateNodeInfo(i, v)
			break
		}
	}

	// if currentReplicated < replicationTarget, execute replication to other node
	if dto.GetCurrentReplicated() < dto.GetReplicationTarget() {

		// find next target replica node
		nextNode, exist := dto.NextReplicaNode()
		if !exist {
			return errors.New("cannot replicate data. No ready data node")
		}

		// execute replication on next node
		err = w.dataNodeService.ReplicateNextNode(ctx, nextNode, dto)
		if err != nil {
			return err
		}

	}

	return nil
}
