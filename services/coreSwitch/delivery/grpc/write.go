package grpc

import (
	coreSwitchProto "github.com/hdkef/hadoop/proto/coreSwitch"
	"github.com/hdkef/hadoop/services/coreSwitch/entity"
)

// Write implements coreSwitch.CoreSwitchServer.
func (h *handler) Write(req *coreSwitchProto.WriteReq, stream coreSwitchProto.CoreSwitch_WriteServer) error {

	dto := &entity.WriteDto{}
	dto.NewFromProto(req)

	progressCh := make(chan uint8)

	err := h.writeUC.Write(stream.Context(), dto, progressCh)

	if err != nil {
		return err
	} else {

		for val := range progressCh {
			stream.Send(&coreSwitchProto.WriteRes{
				Progress: uint32(val),
			})
		}

		return nil
	}
}
