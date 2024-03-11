package impl

import (
	"github.com/hdkef/hadoop/services/dataNode/config"
	"github.com/hdkef/hadoop/services/dataNode/service"
	"github.com/hdkef/hadoop/services/dataNode/usecase"
)

type WriteUsecaseImpl struct {
	cfg             *config.Config
	dataNodeService service.DataNodeService
}

func NewWriteUsecase(cfg *config.Config, dataNodeService service.DataNodeService) usecase.WriteUsecase {

	if cfg == nil {
		panic("nil config")
	}

	return &WriteUsecaseImpl{
		cfg:             cfg,
		dataNodeService: dataNodeService,
	}
}
