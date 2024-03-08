package usecase

import "context"

type WriteUsecase interface {
	Write(ctx context.Context) error
}
