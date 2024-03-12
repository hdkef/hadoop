package entity

import "github.com/google/uuid"

type BlockTarget struct {
	ID      uuid.UUID
	Size    uint64
	NodeIDs []string
}
