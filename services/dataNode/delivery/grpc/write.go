package grpc

import (
	"context"

	"github.com/hdkef/hadoop/pkg/logger"
	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
	"github.com/hdkef/hadoop/services/dataNode/entity"
)

// Create implements dataNode.DataNodeServer.
func (g *handler) Create(ctx context.Context, req *dataNodeProto.CreateReq) (*dataNodeProto.CreateRes, error) {

	// create domain create dto
	createDto := &entity.CreateDto{}
	err := createDto.NewFromProto(req)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	// execute logic
	err = g.writeUC.Create(ctx, createDto)

	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	return &dataNodeProto.CreateRes{}, nil
}
