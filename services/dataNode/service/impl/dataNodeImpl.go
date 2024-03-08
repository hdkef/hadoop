package impl

import (
	"context"
	"fmt"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
	"github.com/hdkef/hadoop/services/dataNode/entity"
	"github.com/hdkef/hadoop/services/dataNode/service"
	"google.golang.org/grpc"
)

type DataNodeService struct {
}

// ReplicateNextNode implements service.DataNodeService.
func (d *DataNodeService) ReplicateNextNode(ctx context.Context, nextNode *pkgEt.NodeInfo, dto *entity.WriteDto) error {
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

	return nil
}

func NewDataNodeService() service.DataNodeService {
	return &DataNodeService{}
}
