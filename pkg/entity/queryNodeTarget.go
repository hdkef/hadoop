package entity

import "github.com/google/uuid"

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
