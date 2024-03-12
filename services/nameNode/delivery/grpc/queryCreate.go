package grpc

import (
	"context"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	nameNodeProto "github.com/hdkef/hadoop/proto/nameNode"
)

// QueryNodeTargetCreate implements nameNode.NameNodeServer.
func (h *handler) QueryNodeTargetCreate(ctx context.Context, req *nameNodeProto.QueryNodeTargetCreateReq) (*nameNodeProto.QueryNodeTarget, error) {

	dto := &pkgEt.CreateReqDto{}
	err := dto.FromProto(req)
	if err != nil {
		return nil, err
	}

	resp, err := h.writeUC.CreateRequest(ctx, dto)
	if err != nil {
		return nil, err
	}

	mappedResponse, err := resp.ToProto()
	if err != nil {
		return nil, err
	}

	return mappedResponse, nil
}
