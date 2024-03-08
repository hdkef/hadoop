package impl

import (
	"context"

	"github.com/hdkef/hadoop/services/coreSwitch/usecase"
)

type WriteUsecaseImpl struct{}

func NewWriteUsecase() usecase.WriteUsecase {
	return &WriteUsecaseImpl{}
}

// Write implements usecase.WriteUsecase.
func (w *WriteUsecaseImpl) Write(ctx context.Context) error {

	// request write to all data nodes concurrently and save blocks metaData to nameNode

	// if success save file metaData to nameNode

	panic("unimplemented")
}
