package impl

import (
	"context"
	"fmt"
	"math/rand"
	"sync"

	"github.com/google/uuid"
	"github.com/hdkef/hadoop/pkg/entity"
	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	pkgSvc "github.com/hdkef/hadoop/pkg/services"
	nameNodeProto "github.com/hdkef/hadoop/proto/nameNode"
	"github.com/hdkef/hadoop/services/client/service"
	"google.golang.org/grpc"
)

type NameNodeService struct {
	nameNodeCache   map[int]*pkgEt.ServiceDiscovery
	serviceRegistry pkgSvc.ServiceRegistry
	mtx             *sync.Mutex
}

// CommitTransaction implements service.NameNodeService.
func (n *NameNodeService) CommitTransaction(ctx context.Context, transactionsID uuid.UUID, isSuccess bool) error {
	// take one nameNode service randomly
	nameNodeSvc := n.nameNodeCache[rand.Intn(len(n.nameNodeCache))+1]
	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", nameNodeSvc.GetAddress(), nameNodeSvc.GetPort()), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := nameNodeProto.NewNameNodeClient(conn)

	// execute commit

	trID, err := transactionsID.MarshalBinary()
	if err != nil {
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
		return err
	}
	return nil
}

// QueryNodeTarget implements service.NameNodeService.
func (n *NameNodeService) QueryNodeTarget(ctx context.Context, dto *pkgEt.CreateReqDto) (*pkgEt.QueryNodeTarget, error) {

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
	nameNodeSvc := n.nameNodeCache[rand.Intn(len(n.nameNodeCache))+1]

	// query nameNode service
	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", nameNodeSvc.GetAddress(), nameNodeSvc.GetPort()), grpc.WithInsecure())
	if err != nil {
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
		return nil, err
	}

	et := &pkgEt.QueryNodeTarget{}

	allBlockID := []uuid.UUID{}
	nodeTarget := []*entity.NodeTarget{}

	for _, v := range resp.GetAllBlockId() {
		b, err := uuid.FromBytes(v)
		if err != nil {
			return nil, err
		}
		allBlockID = append(allBlockID, b)
	}

	for _, v := range resp.GetNodeTarget() {
		newNd := &entity.NodeTarget{}

		bID, err := uuid.FromBytes(v.GetBlockID())
		if err != nil {
			return nil, err
		}

		newNd.SetNodeAddress(v.GetNodeAddress())
		newNd.SetNodeGrpcPort(v.GetNodeGrpcPort())
		newNd.SetNodeID(v.GetNodeID())
		newNd.SetBlockID(bID)
	}

	trId, err := uuid.FromBytes(resp.GetTransactionID())
	if err != nil {
		return nil, err
	}

	iNodeID, err := uuid.FromBytes(resp.GetINodeID())
	if err != nil {
		return nil, err
	}

	et.SetReplicationFactor(resp.GetReplicationFactor())
	et.SetAllBlockID(allBlockID)
	et.SetTransactionID(trId)
	et.SetINodeID(iNodeID)
	et.SetNodeTargets(nodeTarget)

	return et, nil
}

func NewNameNodeService() service.NameNodeService {
	return &NameNodeService{}
}
