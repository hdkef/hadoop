package impl

import (
	"context"
	"fmt"

	nameNodeProto "github.com/hdkef/hadoop/proto/nameNode"
	"github.com/hdkef/hadoop/services/dataNode/config"
	"github.com/hdkef/hadoop/services/dataNode/entity"
	"github.com/hdkef/hadoop/services/dataNode/service"
	"google.golang.org/grpc"
)

type NameNodeService struct {
	cfg *config.Config
}

// UpdateJobQueue implements service.NameNodeService.
func (w *NameNodeService) UpdateJobQueue(ctx context.Context, dto *entity.WriteDto, isSuccess bool) error {
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
		IsSuccess:  isSuccess,
	})
	if err != nil {
		return err
	}
	return nil
}

func NewNameNodeService(cfg *config.Config) service.NameNodeService {

	if cfg == nil {
		panic("nil config")
	}

	return &NameNodeService{
		cfg: cfg,
	}
}
