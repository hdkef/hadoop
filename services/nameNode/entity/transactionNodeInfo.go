package entity

import "github.com/google/uuid"

type TransactionNodeInfo struct {
	nodeID   string
	blockID  uuid.UUID
	address  string
	grpcPort uint32
	fileSize uint64
}

func (t *TransactionNodeInfo) GetNodeID() string {
	return t.nodeID
}

func (t *TransactionNodeInfo) GetBlockID() uuid.UUID {
	return t.blockID
}

func (t *TransactionNodeInfo) GetAddress() string {
	return t.address
}

func (t *TransactionNodeInfo) GetGRPCPort() uint32 {
	return t.grpcPort
}

func (t *TransactionNodeInfo) GetFileSize() uint64 {
	return t.fileSize
}

// Setter methods
func (t *TransactionNodeInfo) SetNodeID(nodeID string) {
	t.nodeID = nodeID
}

func (t *TransactionNodeInfo) SetBlockID(blockID uuid.UUID) {
	t.blockID = blockID
}

func (t *TransactionNodeInfo) SetAddress(address string) {
	t.address = address
}

func (t *TransactionNodeInfo) SetGRPCPort(grpcPort uint32) {
	t.grpcPort = grpcPort
}

func (t *TransactionNodeInfo) SetFileSize(fileSize uint64) {
	t.fileSize = fileSize
}
