package impl

import (
	"context"

	"github.com/hdkef/hadoop/services/nameNode/usecase"
)

type WriteRequestUsecaseImpl struct{}

func NewWriteUsecase() usecase.WriteRequestUsecase {
	return &WriteRequestUsecaseImpl{}
}

// WriteRequest implements usecase.WriteRequestUsecase.
func (w *WriteRequestUsecaseImpl) WriteRequest(ctx context.Context) error {

	// calculate split amount

	// check available name node

	panic("unimplemented")
}
