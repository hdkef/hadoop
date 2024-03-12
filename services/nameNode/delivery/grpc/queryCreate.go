package grpc

import (
	"context"

	nameNodeProto "github.com/hdkef/hadoop/proto/nameNode"
)

// QueryNodeTargetCreate implements nameNode.NameNodeServer.
func (h *handler) QueryNodeTargetCreate(context.Context, *nameNodeProto.QueryNodeTargetCreateReq) (*nameNodeProto.QueryNodeTarget, error) {
	panic("unimplemented")
}
