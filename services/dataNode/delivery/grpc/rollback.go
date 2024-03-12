package grpc

import (
	"context"

	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
)

// Rollback implements dataNode.DataNodeServer.
func (g *handler) Rollback(ctx context.Context, req *dataNodeProto.RollbackReq) (*dataNodeProto.RollbackRes, error) {
	panic("unimplemented")
}
