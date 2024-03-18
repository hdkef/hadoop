package grpc

import (
	"context"

	"github.com/hdkef/hadoop/pkg/logger"
	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
	"github.com/hdkef/hadoop/services/dataNode/entity"
)

// Rollback implements dataNode.DataNodeServer.
func (g *handler) Rollback(ctx context.Context, req *dataNodeProto.RollbackReq) (*dataNodeProto.RollbackRes, error) {

	dto := &entity.RollbackDto{}
	err := dto.FromProto(req)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	err = g.writeUC.RollBack(ctx, dto)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	return &dataNodeProto.RollbackRes{}, nil
}
