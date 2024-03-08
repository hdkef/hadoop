package entity

import (
	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
)

type WriteDto struct {
	inodeID               string
	blockID               string
	blocksData            []byte
	replicationTarget     uint32
	currentReplicated     uint32
	replicationNodeTarget []NodeInfo
	jobQueueID            string
}

// Set methods allow setting individual fields of WriteDto
func (w *WriteDto) SetInodeID(inodeID string) {
	w.inodeID = inodeID
}

func (w *WriteDto) SetBlockID(blockID string) {
	w.blockID = blockID
}

func (w *WriteDto) SetBlocksData(blocksData []byte) {
	w.blocksData = blocksData
}

func (w *WriteDto) SetReplicationTarget(replicationTarget uint32) {
	w.replicationTarget = replicationTarget
}

func (w *WriteDto) SetCurrentReplicated(currentReplicated uint32) {
	w.currentReplicated = currentReplicated
}

func (w *WriteDto) SetReplicationNodeTarget(replicationNodeTarget []NodeInfo) {
	w.replicationNodeTarget = replicationNodeTarget
}

func (w *WriteDto) SetJobQueueID(jobQueueID string) {
	w.jobQueueID = jobQueueID
}

// Get methods allow getting individual fields of WriteDto
func (w *WriteDto) GetInodeID() string {
	return w.inodeID
}

func (w *WriteDto) GetBlockID() string {
	return w.blockID
}

func (w *WriteDto) GetBlocksData() []byte {
	return w.blocksData
}

func (w *WriteDto) GetReplicationTarget() uint32 {
	return w.replicationTarget
}

func (w *WriteDto) GetCurrentReplicated() uint32 {
	return w.currentReplicated
}

func (w *WriteDto) GetReplicationNodeTarget() []NodeInfo {
	return w.replicationNodeTarget
}

func (w *WriteDto) GetJobQueueID() string {
	return w.jobQueueID
}

func (w *WriteDto) IncrementCurrentReplicated() {
	w.currentReplicated++
}

func (w *WriteDto) UpdateNodeInfo(idx int, nodeInfo NodeInfo) {
	if idx < len(w.replicationNodeTarget) {
		w.replicationNodeTarget[idx] = nodeInfo
	}
}

func (w *WriteDto) NextReplicaNode() (NodeInfo, bool) {
	for _, v := range w.replicationNodeTarget {
		if !v.IsSuccess() && !v.IsFailed() {

			return v, true
		}
	}
	return NodeInfo{}, false
}

func (writeDto *WriteDto) NewFromProto(req *dataNodeProto.WriteReq) {

	writeDto.SetInodeID(req.GetINodeID())
	writeDto.SetBlockID(req.GetBlockID())
	writeDto.SetBlocksData(req.GetBlocksData())
	writeDto.SetReplicationTarget(req.GetReplicationTarget())
	writeDto.SetCurrentReplicated(req.GetCurrentReplicated())
	writeDto.SetJobQueueID(req.GetJobQueueID())

	replicationNodeTarget := []NodeInfo{}

	for _, v := range req.GetReplicationNodeTarget() {
		node := NodeInfo{}
		node.SetNodeID(v.GetNodeID())
		node.SetAddress(v.GetAddress())
		node.SetReplicationStatus(ReplicationStatusEnum(v.GetReplicationStatus()))
		node.SetGRPCPort(v.GetGrpcPort())
		replicationNodeTarget = append(replicationNodeTarget, node)
	}

	writeDto.SetReplicationNodeTarget(replicationNodeTarget)
}

func (writeDto *WriteDto) ToProto() *dataNodeProto.WriteReq {

	nodeTarget := []*dataNodeProto.NodeInfo{}

	for _, v := range writeDto.replicationNodeTarget {
		nodeTarget = append(nodeTarget, &dataNodeProto.NodeInfo{
			NodeID:            v.nodeID,
			Address:           v.address,
			GrpcPort:          v.grpcPort,
			ReplicationStatus: v.GetReplicationStatusProto(),
		})
	}

	proto := &dataNodeProto.WriteReq{
		INodeID:               writeDto.inodeID,
		BlockID:               writeDto.blockID,
		BlocksData:            writeDto.blocksData,
		ReplicationTarget:     writeDto.replicationTarget,
		JobQueueID:            writeDto.jobQueueID,
		ReplicationNodeTarget: nodeTarget,
	}

	return proto
}
