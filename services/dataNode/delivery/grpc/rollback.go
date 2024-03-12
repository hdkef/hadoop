package grpc

import (
	"context"

	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
	"github.com/hdkef/hadoop/services/dataNode/entity"
)

// Rollback implements dataNode.DataNodeServer.
func (g *handler) Rollback(ctx context.Context, req *dataNodeProto.RollbackReq) (*dataNodeProto.RollbackRes, error) {

	dto := &entity.RollbackDto{}
	err := dto.FromProto(req)
	if err != nil {
		return nil, err
	}

	err = g.writeUC.RollBack(ctx, dto)
	if err != nil {
		return nil, err
	}

	return &dataNodeProto.RollbackRes{}, nil
}
