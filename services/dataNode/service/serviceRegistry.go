package service

import "github.com/hdkef/hadoop/services/dataNode/config"

type ServiceRegistry interface {
	RegisterDataNode(config *config.Config)
}
