package service

import (
	"context"

	"github.com/hdkef/hadoop/services/client/entity"
)

type DataNodeService interface {
	ReplicateNextNode(ctx context.Context, dto *entity.ReplicateNextNodeDto) error
}
