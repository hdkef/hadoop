package grpc

import (
	"github.com/hdkef/hadoop/services/coreSwitch/config"
	"github.com/hdkef/hadoop/services/coreSwitch/usecase"
	"google.golang.org/grpc"

	coreSwitchProto "github.com/hdkef/hadoop/proto/coreSwitch"
)

type handler struct {
	coreSwitchProto.UnimplementedCoreSwitchServer
	writeUC usecase.WriteUsecase
}

func NewGrpcHandler(cfg *config.Config, writeUC usecase.WriteUsecase) *grpc.Server {

	if cfg == nil {
		panic("nil config")
	}

	grpcServer := grpc.NewServer()
	handler := &handler{
		writeUC: writeUC,
	}

	coreSwitchProto.RegisterCoreSwitchServer(grpcServer, handler)
	return grpcServer
}
