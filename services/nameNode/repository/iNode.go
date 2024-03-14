package repository

import (
	"context"

	"github.com/google/uuid"
	pkgRepoTr "github.com/hdkef/hadoop/pkg/repository/transactionable"
	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type INodeRepo interface {
	Create(ctx context.Context, inode *entity.INode, tx *pkgRepoTr.Transactionable) error
	Get(ctx context.Context, inodeID uuid.UUID, tx *pkgRepoTr.Transactionable) (*entity.INode, error)
	Delete(ctx context.Context, inodeID uuid.UUID, tx *pkgRepoTr.Transactionable) error
}
