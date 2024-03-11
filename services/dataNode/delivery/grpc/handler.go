package grpc

import (
	"github.com/hdkef/hadoop/services/dataNode/config"
	"github.com/hdkef/hadoop/services/dataNode/usecase"
	"google.golang.org/grpc"

	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
)

type handler struct {
	dataNodeProto.UnimplementedDataNodeServer
	writeUC usecase.WriteUsecase
	cfg     *config.Config
}

func NewGrpcHandler(cfg *config.Config, writeUC usecase.WriteUsecase) *grpc.Server {

	if cfg == nil {
		panic("nil config")
	}

	grpcServer := grpc.NewServer()
	handler := &handler{
		writeUC: writeUC,
		cfg:     cfg,
	}

	dataNodeProto.RegisterDataNodeServer(grpcServer, handler)
	return grpcServer
}
