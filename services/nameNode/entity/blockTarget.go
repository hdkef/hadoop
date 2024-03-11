package entity

import "github.com/google/uuid"

type BlockTarget struct {
	ID      uuid.UUID
	Size    uint64
	NodeIDs []string
}

// Getter methods
func (i *INode) GetID() uuid.UUID {
	return i.id
}

func (i *INode) GetBlocks() []*BlockTarget {
	return i.blocks
}

func (i *INode) GetHash() string {
	return i.hash
}

// Setter methods
func (i *INode) SetID(id uuid.UUID) {
	i.id = id
}

func (i *INode) SetBlocks(blocks []*BlockTarget) {
	i.blocks = blocks
}

func (i *INode) SetHash(hash string) {
	i.hash = hash
}
