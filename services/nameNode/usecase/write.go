package usecase

import "context"

type WriteRequestUsecase interface {
	WriteRequest(ctx context.Context) error
}
