package entity

import (
	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
)

type CreateDto struct {
	inodeID               string
	blockID               string
	blocksData            []byte
	replicationTarget     uint32
	currentReplicated     uint32
	replicationNodeTarget []*pkgEt.NodeInfo
}

// Set methods allow setting individual fields of CreateDto
func (w *CreateDto) SetInodeID(inodeID string) {
	w.inodeID = inodeID
}

func (w *CreateDto) SetBlockID(blockID string) {
	w.blockID = blockID
}

func (w *CreateDto) SetBlocksData(blocksData []byte) {
	w.blocksData = blocksData
}

func (w *CreateDto) SetReplicationTarget(replicationTarget uint32) {
	w.replicationTarget = replicationTarget
}

func (w *CreateDto) SetCurrentReplicated(currentReplicated uint32) {
	w.currentReplicated = currentReplicated
}

func (w *CreateDto) SetReplicationNodeTarget(replicationNodeTarget []*pkgEt.NodeInfo) {
	w.replicationNodeTarget = replicationNodeTarget
}

// Get methods allow getting individual fields of CreateDto
func (w *CreateDto) GetInodeID() string {
	return w.inodeID
}

func (w *CreateDto) GetBlockID() string {
	return w.blockID
}

func (w *CreateDto) GetBlocksData() []byte {
	return w.blocksData
}

func (w *CreateDto) GetReplicationTarget() uint32 {
	return w.replicationTarget
}

func (w *CreateDto) GetCurrentReplicated() uint32 {
	return w.currentReplicated
}

func (w *CreateDto) GetReplicationNodeTarget() []*pkgEt.NodeInfo {
	return w.replicationNodeTarget
}

func (w *CreateDto) IncrementCurrentReplicated() {
	w.currentReplicated++
}

func (w *CreateDto) UpdateNodeInfo(idx int, UpdateNodeInfo *pkgEt.NodeInfo) {
	if idx < len(w.replicationNodeTarget) {
		w.replicationNodeTarget[idx] = UpdateNodeInfo
	}
}

func (w *CreateDto) NextReplicaNode() (*pkgEt.NodeInfo, bool) {
	for _, v := range w.replicationNodeTarget {
		if !v.IsSuccess() && !v.IsFailed() {

			return v, true
		}
	}
	return nil, false
}

func (CreateDto *CreateDto) NewFromProto(req *dataNodeProto.CreateReq) {

	CreateDto.SetInodeID(req.GetINodeID())
	CreateDto.SetBlockID(req.GetBlockID())
	CreateDto.SetBlocksData(req.GetBlocksData())
	CreateDto.SetReplicationTarget(req.GetReplicationTarget())
	CreateDto.SetCurrentReplicated(req.GetCurrentReplicated())

	replicationNodeTarget := []*pkgEt.NodeInfo{}

	for _, v := range req.GetReplicationNodeTarget() {
		node := pkgEt.NodeInfo{}
		node.SetNodeID(v.GetNodeID())
		node.SetAddress(v.GetAddress())
		node.SetReplicationStatus(pkgEt.ReplicationStatusEnum(v.GetReplicationStatus()))
		node.SetGRPCPort(v.GetGrpcPort())
		replicationNodeTarget = append(replicationNodeTarget, &node)
	}

	CreateDto.SetReplicationNodeTarget(replicationNodeTarget)
}

func (CreateDto *CreateDto) ToProto() *dataNodeProto.CreateReq {

	nodeTarget := []*dataNodeProto.NodeInfo{}

	for _, v := range CreateDto.replicationNodeTarget {
		nodeTarget = append(nodeTarget, &dataNodeProto.NodeInfo{
			NodeID:            v.GetNodeID(),
			Address:           v.GetAddress(),
			GrpcPort:          v.GetGRPCPort(),
			ReplicationStatus: v.GetReplicationStatusProto(),
		})
	}

	proto := &dataNodeProto.CreateReq{
		INodeID:               CreateDto.inodeID,
		BlockID:               CreateDto.blockID,
		BlocksData:            CreateDto.blocksData,
		ReplicationTarget:     CreateDto.replicationTarget,
		ReplicationNodeTarget: nodeTarget,
	}

	return proto
}
