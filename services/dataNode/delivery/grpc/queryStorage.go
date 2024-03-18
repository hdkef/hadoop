package grpc

import (
	"context"
	"fmt"
	"syscall"

	"github.com/hdkef/hadoop/pkg/logger"
	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
)

// QueryStorage implements dataNode.DataNodeServer.
func (g *handler) QueryStorage(ctx context.Context, req *dataNodeProto.QueryStorageReq) (*dataNodeProto.QueryStorageRes, error) {

	var stat syscall.Statfs_t

	err := syscall.Statfs(g.cfg.StorageRoot, &stat)
	if err != nil {
		logger.LogError(err)
		fmt.Println("Error:", err)
		return nil, err
	}

	// Calculate total size in bytes
	total := stat.Blocks * uint64(stat.Bsize)
	// Calculate used size in bytes
	used := (stat.Blocks - stat.Bfree) * uint64(stat.Bsize)
	return &dataNodeProto.QueryStorageRes{
		TotalStorage:      total,
		ActualUsedStorage: used,
	}, nil
}
