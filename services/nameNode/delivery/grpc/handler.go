package grpc

import (
	"github.com/hdkef/hadoop/services/nameNode/config"
	"github.com/hdkef/hadoop/services/nameNode/usecase"
	"google.golang.org/grpc"

	nameNodeProto "github.com/hdkef/hadoop/proto/nameNode"
)

type handler struct {
	nameNodeProto.UnimplementedNameNodeServer
	writeUC usecase.WriteRequestUsecase
	cfg     *config.Config
}

func NewGrpcHandler(cfg *config.Config, writeUC *usecase.WriteRequestUsecase) *grpc.Server {

	if cfg == nil {
		panic("nil config")
	}

	if writeUC == nil {
		panic("writeUC is nil")
	}

	grpcServer := grpc.NewServer()
	handler := &handler{
		writeUC: *writeUC,
		cfg:     cfg,
	}

	nameNodeProto.RegisterNameNodeServer(grpcServer, handler)
	return grpcServer
}
