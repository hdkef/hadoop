package grpc

import (
	"github.com/hdkef/hadoop/services/client/config"
	"github.com/hdkef/hadoop/services/client/usecase"
	"google.golang.org/grpc"

	clientProto "github.com/hdkef/hadoop/proto/client"
)

type handler struct {
	clientProto.UnimplementedClientServer
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

	clientProto.RegisterClientServer(grpcServer, handler)
	return grpcServer
}
