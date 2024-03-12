package service

import (
	"context"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
)

type NameNodeService interface {
	QueryNodeTarget(ctx context.Context, dto *pkgEt.CreateReqDto) (*pkgEt.QueryNodeTarget, error)
}
