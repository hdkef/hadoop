package entity

import (
	pkgEt "github.com/hdkef/hadoop/pkg/entity"
)

type ReplicateNextNodeDto struct {
	iNodeID               string
	blockID               string
	blocksData            []byte
	replicationTarget     uint32
	currentReplicated     uint32
	replicationNodeTarget []*pkgEt.NodeInfo
	jobQueueID            string
	nextNode              *pkgEt.NodeInfo
}

func (r *ReplicateNextNodeDto) GetINodeID() string {
	return r.iNodeID
}

func (r *ReplicateNextNodeDto) SetINodeID(inodeID string) {
	r.iNodeID = inodeID
}

func (r *ReplicateNextNodeDto) GetBlockID() string {
	return r.blockID
}

func (r *ReplicateNextNodeDto) SetBlockID(blockID string) {
	r.blockID = blockID
}

func (r *ReplicateNextNodeDto) GetBlocksData() []byte {
	return r.blocksData
}

func (r *ReplicateNextNodeDto) SetBlocksData(blocksData []byte) {
	r.blocksData = blocksData
}

func (r *ReplicateNextNodeDto) GetReplicationTarget() uint32 {
	return r.replicationTarget
}

func (r *ReplicateNextNodeDto) SetReplicationTarget(replicationTarget uint32) {
	r.replicationTarget = replicationTarget
}

func (r *ReplicateNextNodeDto) GetCurrentReplicated() uint32 {
	return r.currentReplicated
}

func (r *ReplicateNextNodeDto) SetCurrentReplicated(currentReplicated uint32) {
	r.currentReplicated = currentReplicated
}

func (r *ReplicateNextNodeDto) GetReplicationNodeTarget() []*pkgEt.NodeInfo {
	return r.replicationNodeTarget
}

func (r *ReplicateNextNodeDto) SetReplicationNodeTarget(replicationNodeTarget []*pkgEt.NodeInfo) {
	r.replicationNodeTarget = replicationNodeTarget
}

func (r *ReplicateNextNodeDto) GetJobQueueID() string {
	return r.jobQueueID
}

func (r *ReplicateNextNodeDto) SetJobQueueID(jobQueueID string) {
	r.jobQueueID = jobQueueID
}

func (r *ReplicateNextNodeDto) GetNextNode() *pkgEt.NodeInfo {
	return r.nextNode
}

func (r *ReplicateNextNodeDto) SetNextNode(nextNode *pkgEt.NodeInfo) {
	r.nextNode = nextNode
}
