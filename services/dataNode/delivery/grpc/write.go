package grpc

import (
	"context"

	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
	"github.com/hdkef/hadoop/services/dataNode/entity"
)

// Write implements dataNode.DataNodeServer.
func (g *handler) Write(ctx context.Context, req *dataNodeProto.CreateReq) (*dataNodeProto.CreateRes, error) {

	// create domain create dto
	createDto := &entity.CreateDto{}
	createDto.NewFromProto(req)

	// execute logic
	err := g.writeUC.Create(ctx, createDto)

	if err != nil {
		return nil, err
	}

	return &dataNodeProto.CreateRes{}, nil
}
