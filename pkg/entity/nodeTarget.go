package entity

import "github.com/google/uuid"

type NodeTarget struct {
	nodeID       string
	nodeAddress  string
	nodeGrpcPort uint32
	blockID      uuid.UUID
}

func (n *NodeTarget) GetNodeID() string {
	return n.nodeID
}

func (n *NodeTarget) SetNodeID(nodeID string) {
	n.nodeID = nodeID
}

func (n *NodeTarget) GetNodeAddress() string {
	return n.nodeAddress
}

func (n *NodeTarget) SetNodeAddress(nodeAddress string) {
	n.nodeAddress = nodeAddress
}

func (n *NodeTarget) GetNodeGrpcPort() uint32 {
	return n.nodeGrpcPort
}

func (n *NodeTarget) SetNodeGrpcPort(nodePort uint32) {
	n.nodeGrpcPort = nodePort
}

func (n *NodeTarget) GetBlockID() uuid.UUID {
	return n.blockID
}

func (n *NodeTarget) SetBlockID(blockID uuid.UUID) {
	n.blockID = blockID
}
