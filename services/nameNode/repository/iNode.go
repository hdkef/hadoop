package repository

import (
	"context"

	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type INodeRepo interface {
	Create(ctx context.Context, inode *entity.INode) error
	Get(ctx context.Context, inodeID string) (*entity.INode, error)
	Delete(ctx context.Context, inodeID string) error
}
