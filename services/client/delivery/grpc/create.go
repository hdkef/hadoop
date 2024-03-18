package grpc

import (
	"github.com/hdkef/hadoop/pkg/logger"
	clientProto "github.com/hdkef/hadoop/proto/client"
	"github.com/hdkef/hadoop/services/client/entity"
)

// Create implements client.ClientServer.
func (h *handler) Create(req *clientProto.CreateReq, stream clientProto.Client_CreateServer) error {

	dto := &entity.CreateDto{}
	err := dto.NewFromProto(req)
	if err != nil {
		logger.LogError(err)
		return err
	}

	progressCh := make(chan entity.CreateStreamRes)

	go h.writeUC.Create(stream.Context(), dto, progressCh)

	for val := range progressCh {

		if val.IsError() {
			logger.LogError(val.GetError())
			return val.GetError()
		}

		stream.Send(&clientProto.CreateRes{
			Progress: uint32(val.GetProgress()),
		})
	}

	return nil
}
