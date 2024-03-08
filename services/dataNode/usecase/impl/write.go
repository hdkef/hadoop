package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/hdkef/hadoop/services/dataNode/config"
	"github.com/hdkef/hadoop/services/dataNode/entity"
	"github.com/hdkef/hadoop/services/dataNode/usecase"
	"google.golang.org/grpc"

	pkgRepo "github.com/hdkef/hadoop/pkg/repository"
	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
	nameNodeProto "github.com/hdkef/hadoop/proto/nameNode"
)

type WriteUsecaseImpl struct {
	cfg    *config.Config
	kvRepo pkgRepo.KeyValueRepository
}

func NewWriteUsecase(cfg *config.Config, kvRepo pkgRepo.KeyValueRepository) usecase.WriteUsecase {

	if cfg == nil {
		panic("nil config")
	}

	return &WriteUsecaseImpl{
		cfg:    cfg,
		kvRepo: kvRepo,
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
	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", w.cfg.NameNodeAddress, w.cfg.NameNodePort), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := nameNodeProto.NewNameNodeClient(conn)
	_, err = client.UpdateJobQueue(ctx, &nameNodeProto.UpdateJobQueueReq{
		JobQueueID: dto.GetJobQueueID(),
		INodeID:    dto.GetInodeID(),
		BlockID:    dto.GetBlockID(),
		NodeID:     w.cfg.NodeId,
		UpdateType: nameNodeProto.UpdateJobQueueReq_REPLICATE_STATUS,
	})
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
		conn, err := grpc.Dial(fmt.Sprintf("%v:%d", nextNode.GetAddress(), nextNode.GetGRPCPort()), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := dataNodeProto.NewDataNodeClient(conn)
		_, err = client.Write(ctx, dto.ToProto())
		if err != nil {
			return err
		}

	}

	return nil
}
