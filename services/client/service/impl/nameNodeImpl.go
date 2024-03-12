package impl

import (
	"context"
	"sync"

	"github.com/hdkef/hadoop/pkg/entity"
	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	pkgSvc "github.com/hdkef/hadoop/pkg/services"
	"github.com/hdkef/hadoop/services/client/service"
)

type NameNodeService struct {
	nameNodeCache   map[int]*pkgEt.ServiceDiscovery
	serviceRegistry pkgSvc.ServiceRegistry
	mtx             *sync.Mutex
}

// QueryNodeTarget implements service.NameNodeService.
func (n *NameNodeService) QueryNodeTarget(ctx context.Context, dto *entity.CreateReqDto) (*entity.QueryNodeTarget, error) {

	// if nameNode empty, query service registry
	if len(n.nameNodeCache) == 0 {
		svd, err := n.serviceRegistry.GetAll(ctx, "nameNode", "")
		if err != nil {
			return nil, err
		}

		n.mtx.Lock()
		for i, v := range svd {
			n.nameNodeCache[i] = v
		}
		n.mtx.Unlock()
	}

	// take one nameNode service randomly
	// nameNodeSvc := n.nameNodeCache[rand.Intn(len(n.nameNodeCache))+1]

	// TODO
	// query nameNode service
	return nil, nil
}

func NewNameNodeService() service.NameNodeService {
	return &NameNodeService{}
}
