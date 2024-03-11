package impl

import (
	"context"
	"fmt"

	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
	"github.com/hdkef/hadoop/services/client/entity"
	"github.com/hdkef/hadoop/services/client/service"
	"google.golang.org/grpc"
)

type DataNodeService struct {
}

// ReplicateNextNode implements service.DataNodeService.
func (d *DataNodeService) ReplicateNextNode(ctx context.Context, dto *entity.ReplicateNextNodeDto) error {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", dto.GetNextNode().GetAddress(), dto.GetNextNode().GetGRPCPort()), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := dataNodeProto.NewDataNodeClient(conn)

	nodeTarget := []*dataNodeProto.NodeInfo{}

	for _, v := range dto.GetReplicationNodeTarget() {
		nodeTarget = append(nodeTarget, &dataNodeProto.NodeInfo{
			NodeID:            v.GetNodeID(),
			Address:           v.GetAddress(),
			GrpcPort:          v.GetGRPCPort(),
			ReplicationStatus: v.GetReplicationStatusProto(),
		})
	}

	_, err = client.Create(ctx, &dataNodeProto.CreateReq{
		INodeID:               dto.GetINodeID().String(),
		BlockID:               dto.GetBlockID().String(),
		BlocksData:            dto.GetBlocksData(),
		ReplicationTarget:     dto.GetReplicationTarget(),
		CurrentReplicated:     dto.GetCurrentReplicated(),
		ReplicationNodeTarget: nodeTarget,
	})
	if err != nil {
		return err
	}

	return nil
}

func NewDataNodeService() service.DataNodeService {
	return &DataNodeService{}
}
