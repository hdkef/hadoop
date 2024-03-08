package entity

import dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"

type ReplicationStatusEnum uint8

const REPLICATION_STATUS_SUCCESS ReplicationStatusEnum = 0
const REPLICATION_STATUS_FAILED ReplicationStatusEnum = 1
const REPLICATION_STATUS_PENDING ReplicationStatusEnum = 2

type NodeInfo struct {
	nodeID            string
	address           string
	grpcPort          uint32
	replicationStatus ReplicationStatusEnum
}

// Set methods allow setting individual fields of NodeInfo
func (n *NodeInfo) SetNodeID(nodeID string) {
	n.nodeID = nodeID
}

func (n *NodeInfo) SetAddress(address string) {
	n.address = address
}

func (n *NodeInfo) SetGRPCPort(grpcPort uint32) {
	n.grpcPort = grpcPort
}

func (n *NodeInfo) SetReplicationStatusSuccess() {
	n.replicationStatus = REPLICATION_STATUS_SUCCESS
}

func (n *NodeInfo) SetReplicationStatus(status ReplicationStatusEnum) {
	n.replicationStatus = status
}

// Get methods allow getting individual fields of NodeInfo
func (n *NodeInfo) GetNodeID() string {
	return n.nodeID
}

func (n *NodeInfo) GetAddress() string {
	return n.address
}

func (n *NodeInfo) GetGRPCPort() uint32 {
	return n.grpcPort
}

func (n *NodeInfo) GetReplicationStatus() ReplicationStatusEnum {
	return n.replicationStatus
}

func (n *NodeInfo) IsSuccess() bool {
	return n.replicationStatus == REPLICATION_STATUS_SUCCESS
}

func (n *NodeInfo) IsFailed() bool {
	return n.replicationStatus == REPLICATION_STATUS_FAILED
}

func (n *NodeInfo) GetReplicationStatusProto() dataNodeProto.NodeInfo_NodeReplicationStatus {

	switch n.replicationStatus {
	case REPLICATION_STATUS_SUCCESS:
		return dataNodeProto.NodeInfo_SUCCESS
	case REPLICATION_STATUS_FAILED:
		return dataNodeProto.NodeInfo_FAILED
	default:
		return dataNodeProto.NodeInfo_PENDING
	}

}
