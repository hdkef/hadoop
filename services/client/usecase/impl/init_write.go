package impl

import (
	"github.com/hdkef/hadoop/services/client/service"
	"github.com/hdkef/hadoop/services/client/usecase"
)

type WriteUsecaseImpl struct {
	dataNodeService service.DataNodeService
}

func NewWriteUsecase(dataNodeService service.DataNodeService) usecase.WriteUsecase {
	return &WriteUsecaseImpl{
		dataNodeService: dataNodeService,
	}
}
