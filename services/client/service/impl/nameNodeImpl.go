package impl

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	"github.com/hdkef/hadoop/pkg/logger"
	pkgSvc "github.com/hdkef/hadoop/pkg/services"
	pkgSvcImpl "github.com/hdkef/hadoop/pkg/services/impl"
	nameNodeProto "github.com/hdkef/hadoop/proto/nameNode"
	"github.com/hdkef/hadoop/services/client/service"
	"google.golang.org/grpc"
)

type NameNodeServiceDto struct {
	NameNodeCache   map[int]*pkgEt.ServiceDiscovery
	ServiceRegistry *pkgSvc.ServiceRegistry
	Mtx             *sync.Mutex
}

type NameNodeService struct {
	nameNodeCache   map[int]*pkgEt.ServiceDiscovery
	serviceRegistry pkgSvc.ServiceRegistry
	mtx             *sync.Mutex
	loadBalancer    pkgSvc.RoundRobinLoadBalancer
}

// CommitTransaction implements service.NameNodeService.
func (n *NameNodeService) CommitTransaction(ctx context.Context, transactionsID uuid.UUID, isSuccess bool) error {
	// take one nameNode service randomly
	n.mtx.Lock()

	ptr, err := n.loadBalancer.GetNextPtr(len(n.nameNodeCache))
	if err != nil {
		n.mtx.Unlock()
		return err
	}
	nameNodeSvc := n.nameNodeCache[ptr]

	n.mtx.Unlock()

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", nameNodeSvc.GetAddress(), nameNodeSvc.GetPort()), grpc.WithInsecure())
	if err != nil {
		logger.LogError(err)
		return err
	}
	defer conn.Close()

	client := nameNodeProto.NewNameNodeClient(conn)

	// execute commit

	trID, err := transactionsID.MarshalBinary()
	if err != nil {
		logger.LogError(err)
		return err
	}

	status := nameNodeProto.CommitTransactionsReq_FAILED
	if isSuccess {
		status = nameNodeProto.CommitTransactionsReq_SUCCESS
	}

	_, err = client.CommitTransactions(ctx, &nameNodeProto.CommitTransactionsReq{
		TransactionID: trID,
		Status:        status,
	})
	if err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

// QueryNodeTarget implements service.NameNodeService.
func (n *NameNodeService) QueryNodeTarget(ctx context.Context, dto *pkgEt.CreateReqDto) (*pkgEt.QueryNodeTarget, error) {

	// if nameNode empty, query service registry
	n.mtx.Lock()
	if len(n.nameNodeCache) == 0 {
		n.mtx.Unlock()
		svd, err := n.serviceRegistry.GetAll(ctx, "nameNode", "")
		if err != nil {
			logger.LogError(err)
			return nil, err
		}

		n.mtx.Lock()
		for i, v := range svd {
			n.nameNodeCache[i] = v
		}
	}

	// take one nameNode service randomly
	ptr, err := n.loadBalancer.GetNextPtr(len(n.nameNodeCache))
	if err != nil {
		n.mtx.Unlock()
		return nil, err
	}
	nameNodeSvc := n.nameNodeCache[ptr]

	n.mtx.Unlock()

	// query nameNode service
	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", nameNodeSvc.GetAddress(), nameNodeSvc.GetPort()), grpc.WithInsecure())
	if err != nil {
		logger.LogError(err)
		return nil, err
	}
	defer conn.Close()

	client := nameNodeProto.NewNameNodeClient(conn)

	resp, err := client.QueryNodeTargetCreate(ctx, &nameNodeProto.QueryNodeTargetCreateReq{
		ReplicationTarget: dto.GetReplicationTarget(),
		BlockSplitTarget:  dto.GetBlockSplitTarget(),
		FileSize:          dto.GetFileSize(),
		LeaseTimeInSec:    dto.GetLeaseTimeInSec(),
		Path:              dto.GetPath(),
		Hash:              dto.GetHash(),
	})
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	et := &pkgEt.QueryNodeTarget{}

	allBlockID := []uuid.UUID{}
	nodeTarget := []*pkgEt.NodeTarget{}

	for _, v := range resp.GetAllBlockId() {
		b, err := uuid.FromBytes(v)
		if err != nil {
			logger.LogError(err)
			return nil, err
		}
		allBlockID = append(allBlockID, b)
	}

	for _, v := range resp.GetNodeTarget() {
		newNd := &pkgEt.NodeTarget{}

		bID, err := uuid.FromBytes(v.GetBlockID())
		if err != nil {
			logger.LogError(err)
			return nil, err
		}

		newNd.SetNodeAddress(v.GetNodeAddress())
		newNd.SetNodeGrpcPort(v.GetNodeGrpcPort())
		newNd.SetNodeID(v.GetNodeID())
		newNd.SetBlockID(bID)
		nodeTarget = append(nodeTarget, newNd)
	}

	trId, err := uuid.FromBytes(resp.GetTransactionID())
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	iNodeID, err := uuid.FromBytes(resp.GetINodeID())
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	et.SetReplicationFactor(resp.GetReplicationFactor())
	et.SetAllBlockID(allBlockID)
	et.SetTransactionID(trId)
	et.SetINodeID(iNodeID)
	et.SetNodeTargets(nodeTarget)

	return et, nil
}

func NewNameNodeService(dto *NameNodeServiceDto) service.NameNodeService {

	if dto.NameNodeCache == nil {
		panic("nameNodeCache nil")
	}

	if dto.ServiceRegistry == nil {
		panic("serviceRegistry nil")
	}

	if dto.Mtx == nil {
		panic("mtx nil")
	}

	return &NameNodeService{
		nameNodeCache:   dto.NameNodeCache,
		serviceRegistry: *dto.ServiceRegistry,
		mtx:             dto.Mtx,
		loadBalancer:    pkgSvcImpl.NewLoadBalancer(&sync.Mutex{}),
	}
}
