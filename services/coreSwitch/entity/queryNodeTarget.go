package entity

type QueryNodeTarget struct {
	replicationFactor uint32
	allBlockID        []string
	jobQueueID        string
	iNodeID           string
	nodeTargets       []*NodeTarget
}

func (q *QueryNodeTarget) GetINodeID() string {
	return q.iNodeID
}

func (q *QueryNodeTarget) GetJobQueueID() string {
	return q.jobQueueID
}

func (q *QueryNodeTarget) GetBlockID(idx int) string {
	return q.allBlockID[idx]
}

func (q *QueryNodeTarget) GetNodeTarget(blockID string) []*NodeTarget {

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
