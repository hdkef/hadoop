package impl

import (
	"context"
	"fmt"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
	"google.golang.org/grpc"

	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/service"
)

type DataNodeService struct{}

// Rollback implements service.DataNodeService.
func (d *DataNodeService) Rollback(ctx context.Context, dto *entity.RollbackDto) error {

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", dto.GetNodeAddress(), dto.GetNodePort()), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := dataNodeProto.NewDataNodeClient(conn)

	_, err = client.Rollback(ctx, &dataNodeProto.RollbackReq{
		INodeID: dto.GetINodeID().String(),
		BlockID: dto.GetBlockID().String(),
	})
	if err != nil {
		return err
	}

	return nil
}

// QueryStorage implements service.DataNodeService.
func (d *DataNodeService) QueryStorage(ctx context.Context, svd *pkgEt.ServiceDiscovery) (*entity.NodeStorage, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", svd.GetAddress(), svd.GetPort()), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := dataNodeProto.NewDataNodeClient(conn)

	resp, err := client.QueryStorage(ctx, &dataNodeProto.QueryStorageReq{})
	if err != nil {
		return nil, err
	}

	nd := &entity.NodeStorage{}
	nd.SetActualUsedStorage(resp.GetActualUsedStorage())
	nd.SetTotalStorage(resp.GetTotalStorage())
	nd.SetNodeID(svd.GetID())

	return nd, nil
}

func NewDataNodeService() service.DataNodeService {
	return &DataNodeService{}
}
