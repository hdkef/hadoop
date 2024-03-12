package entity

import (
	"github.com/google/uuid"
	nameNodeProto "github.com/hdkef/hadoop/proto/nameNode"
)

type QueryNodeTarget struct {
	replicationFactor uint32
	allBlockID        []uuid.UUID
	transactionID     uuid.UUID
	iNodeID           uuid.UUID
	nodeTargets       []*NodeTarget
}

func (q *QueryNodeTarget) GetINodeID() uuid.UUID {
	return q.iNodeID
}

func (q *QueryNodeTarget) GetTransactionID() uuid.UUID {
	return q.transactionID
}

func (q *QueryNodeTarget) GetBlockID(idx int) uuid.UUID {
	return q.allBlockID[idx]
}

func (q *QueryNodeTarget) GetNodeTarget(blockID uuid.UUID) []*NodeTarget {

	allNodeTarget := []*NodeTarget{}

	for _, v := range q.nodeTargets {
		if v.blockID == blockID {
			allNodeTarget = append(allNodeTarget, v)
		}
	}

	return allNodeTarget

}

func (q *QueryNodeTarget) GetTotalBlock() int {
	return len(q.allBlockID)
}

func (q *QueryNodeTarget) SetReplicationFactor(replicationFactor uint32) {
	q.replicationFactor = replicationFactor
}

func (q *QueryNodeTarget) SetAllBlockID(allBlockID []uuid.UUID) {
	q.allBlockID = allBlockID
}

func (q *QueryNodeTarget) SetTransactionID(transactionID uuid.UUID) {
	q.transactionID = transactionID
}

func (q *QueryNodeTarget) SetINodeID(iNodeID uuid.UUID) {
	q.iNodeID = iNodeID
}

func (q *QueryNodeTarget) SetNodeTargets(nodeTargets []*NodeTarget) {
	q.nodeTargets = nodeTargets
}

func (q *QueryNodeTarget) ToProto() (*nameNodeProto.QueryNodeTarget, error) {

	nodeTarget := []*nameNodeProto.NodeTarget{}

	for _, v := range q.nodeTargets {

		bId, err := v.blockID.MarshalBinary()
		if err != nil {
			return nil, err
		}

		nodeTarget = append(nodeTarget, &nameNodeProto.NodeTarget{
			NodeID:       v.nodeID,
			NodeAddress:  v.nodeAddress,
			NodeGrpcPort: v.nodeGrpcPort,
			BlockID:      bId,
		})
	}

	allBlocksID := [][]byte{}

	for _, v := range q.allBlockID {
		bId, err := v.MarshalBinary()
		if err != nil {
			return nil, err
		}
		allBlocksID = append(allBlocksID, bId)
	}

	trId, err := q.transactionID.MarshalBinary()
	if err != nil {
		return nil, err
	}

	iNodeId, err := q.iNodeID.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return &nameNodeProto.QueryNodeTarget{
		ReplicationFactor: q.replicationFactor,
		AllBlockId:        allBlocksID,
		TransactionID:     trId,
		INodeID:           iNodeId,
		NodeTarget:        nodeTarget,
	}, nil
}
