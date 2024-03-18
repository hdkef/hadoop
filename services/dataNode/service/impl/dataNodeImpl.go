package impl

import (
	"context"
	"fmt"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	"github.com/hdkef/hadoop/pkg/logger"
	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
	"github.com/hdkef/hadoop/services/dataNode/entity"
	"github.com/hdkef/hadoop/services/dataNode/service"
	"google.golang.org/grpc"
)

type DataNodeService struct {
}

// ReplicateNextNode implements service.DataNodeService.
func (d *DataNodeService) ReplicateNextNode(ctx context.Context, nextNode *pkgEt.NodeInfo, dto *entity.CreateDto) error {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", nextNode.GetAddress(), nextNode.GetGRPCPort()), grpc.WithInsecure())
	if err != nil {
		logger.LogError(err)
		return err
	}
	defer conn.Close()

	client := dataNodeProto.NewDataNodeClient(conn)

	dtoProto, err := dto.ToProto()
	if err != nil {
		logger.LogError(err)
		return err
	}

	_, err = client.Create(ctx, dtoProto)
	if err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func NewDataNodeService() service.DataNodeService {
	return &DataNodeService{}
}
