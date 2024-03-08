package impl

import (
	"context"

	"github.com/hdkef/hadoop/services/dataNode/usecase"
)

type WriteUsecaseImpl struct{}

func NewWriteUsecase() usecase.WriteUsecase {
	return &WriteUsecaseImpl{}
}

// Write implements usecase.WriteUsecase.
func (w *WriteUsecaseImpl) Write(ctx context.Context) error {

	// create new k,v store

	// write to storage

	panic("unimplemented")
}
