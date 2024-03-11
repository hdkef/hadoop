package entity

import "github.com/google/uuid"

type INode struct {
	id     uuid.UUID
	blocks []*BlockTarget
	hash   string
}
