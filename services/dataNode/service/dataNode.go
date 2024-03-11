package service

import (
	"context"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	"github.com/hdkef/hadoop/services/dataNode/entity"
)

type DataNodeService interface {
	ReplicateNextNode(ctx context.Context, nextNode *pkgEt.NodeInfo, dto *entity.CreateDto) error
}
