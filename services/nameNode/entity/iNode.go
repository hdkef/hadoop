package entity

import "github.com/google/uuid"

type INode struct {
	id          uuid.UUID
	blocks      []*BlockTarget
	allBlockIDs []uuid.UUID
}

// Getter methods
func (i *INode) GetID() uuid.UUID {
	return i.id
}

func (i *INode) GetBlocks() []*BlockTarget {
	return i.blocks
}

// Setter methods
func (i *INode) SetID(id uuid.UUID) {
	i.id = id
}

func (i *INode) SetBlocks(blocks []*BlockTarget) {
	i.blocks = blocks
}

func (t *INode) SetAllBlockIds(val []uuid.UUID) {
	t.allBlockIDs = val
}

func (t *INode) GetAllBlockIds() []uuid.UUID {
	return t.allBlockIDs
}
