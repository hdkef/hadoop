package impl

import (
	"github.com/hdkef/hadoop/services/client/service"
	"github.com/hdkef/hadoop/services/client/usecase"
)

type WriteUsecaseImpl struct {
	dataNodeService service.DataNodeService
	nameNodeService service.NameNodeService
}

func NewWriteUsecase(dataNodeService *service.DataNodeService, nameNodeService *service.NameNodeService) usecase.WriteUsecase {

	if dataNodeService == nil {
		panic("dataNodeService is nil")
	}

	if nameNodeService == nil {
		panic("nameNodeService is nil")
	}

	return &WriteUsecaseImpl{
		dataNodeService: *dataNodeService,
		nameNodeService: *nameNodeService,
	}
}
