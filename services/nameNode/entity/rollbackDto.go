package entity

import (
	"github.com/google/uuid"
	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
)

type RollbackDto struct {
	nodeID      string
	nodeAddress string
	nodePort    uint32
	iNodeID     uuid.UUID
	blockID     uuid.UUID
}

// Getter methods
func (r *RollbackDto) GetNodeID() string {
	return r.nodeID
}

func (r *RollbackDto) GetNodeAddress() string {
	return r.nodeAddress
}

func (r *RollbackDto) GetNodePort() uint32 {
	return r.nodePort
}

func (r *RollbackDto) GetINodeID() uuid.UUID {
	return r.iNodeID
}

func (r *RollbackDto) GetBlockID() uuid.UUID {
	return r.blockID
}

// Setter methods
func (r *RollbackDto) SetNodeID(nodeID string) {
	r.nodeID = nodeID
}

func (r *RollbackDto) SetNodeAddress(nodeAddress string) {
	r.nodeAddress = nodeAddress
}

func (r *RollbackDto) SetNodePort(nodePort uint32) {
	r.nodePort = nodePort
}

func (r *RollbackDto) SetINodeID(iNodeID uuid.UUID) {
	r.iNodeID = iNodeID
}

func (r *RollbackDto) SetBlockID(blockID uuid.UUID) {
	r.blockID = blockID
}

func (r *RollbackDto) ToProto() (*dataNodeProto.RollbackReq, error) {

	iNdID, err := r.iNodeID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	bID, err := r.blockID.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return &dataNodeProto.RollbackReq{
		INodeID: iNdID,
		BlockID: bID,
	}, nil
}
