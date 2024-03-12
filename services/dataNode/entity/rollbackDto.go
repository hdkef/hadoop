package entity

import "github.com/google/uuid"

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
