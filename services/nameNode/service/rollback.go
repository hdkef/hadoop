package service

import (
	"context"

	"github.com/hdkef/hadoop/services/nameNode/entity"
)

type RollbackService interface {
	Rollback(ctx context.Context, tx *entity.Transactions) error
}
