package service

import (
	"context"

	"github.com/hdkef/hadoop/services/dataNode/entity"
)

type NameNodeService interface {
	UpdateJobQueue(ctx context.Context, dto *entity.WriteDto, isSuccess bool) error
}
