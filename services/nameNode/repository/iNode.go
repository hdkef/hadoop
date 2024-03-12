package repository

import (
	"context"

	pkgRepo "github.com/hdkef/hadoop/pkg/repository"
	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type INodeRepo interface {
	Create(ctx context.Context, inode *entity.INode, tx pkgRepo.Transactionable) error
	Get(ctx context.Context, inodeID string, tx pkgRepo.Transactionable) (*entity.INode, error)
	Delete(ctx context.Context, inodeID string, tx pkgRepo.Transactionable) error
}
