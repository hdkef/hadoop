package entity

import (
	"github.com/google/uuid"
	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
)

type RollbackDto struct {
	iNodeID uuid.UUID
	blockID uuid.UUID
}

func (r *RollbackDto) GetINodeID() uuid.UUID {
	return r.iNodeID
}

func (r *RollbackDto) GetBlockID() uuid.UUID {
	return r.blockID
}

func (r *RollbackDto) SetINodeID(iNodeID uuid.UUID) {
	r.iNodeID = iNodeID
}

func (r *RollbackDto) SetBlockID(blockID uuid.UUID) {
	r.blockID = blockID
}

func (r *RollbackDto) FromProto(req *dataNodeProto.RollbackReq) error {
	bId, err := uuid.FromBytes(req.GetBlockID())
	if err != nil {
		return err
	}
	iNodeId, err := uuid.FromBytes(req.GetINodeID())
	if err != nil {
		return err
	}
	r.blockID = bId
	r.iNodeID = iNodeId

	return nil
}
