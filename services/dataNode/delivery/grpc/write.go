package grpc

import (
	"context"

	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
	"github.com/hdkef/hadoop/services/dataNode/entity"
)

// Write implements dataNode.DataNodeServer.
func (g *handler) Write(ctx context.Context, req *dataNodeProto.WriteReq) (*dataNodeProto.WriteRes, error) {

	// create domain write dto
	writeDto := &entity.WriteDto{}
	writeDto.NewFromProto(req)

	// execute logic
	err := g.writeUC.Write(ctx, writeDto)

	if err != nil {
		return nil, err
	}

	return &dataNodeProto.WriteRes{}, nil
}
