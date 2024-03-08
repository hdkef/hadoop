package impl

import (
	"context"

	"github.com/hdkef/hadoop/services/client/entity"
	"github.com/hdkef/hadoop/services/client/usecase"
)

type WriteUsecaseImpl struct {
}

func NewWriteUsecaseImpl() usecase.WriteUsecase {
	return &WriteUsecaseImpl{}
}

// Write implements usecase.WriteUsecase.
func (w *WriteUsecaseImpl) Write(ctx context.Context, dto entity.WriteRequestDto) error {

	// request write to name node

	// request write to core switching

	return nil
}
