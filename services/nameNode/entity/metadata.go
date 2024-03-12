package entity

import "github.com/google/uuid"

const (
	METADATA_TYPE_DIR  = 1
	METADATA_TYPE_FILE = 2
)

type MetadataType uint8

type Metadata struct {
	parentPath string
	path       string
	mtype      MetadataType
	iNodeID    uuid.UUID
	hash       string
}

// Getter methods
func (m *Metadata) GetParentPath() string {
	return m.parentPath
}

func (m *Metadata) GetPath() string {
	return m.path
}

func (m *Metadata) GetType() MetadataType {
	return m.mtype
}

func (m *Metadata) GetINodeID() uuid.UUID {
	return m.iNodeID
}

// Setter methods
func (m *Metadata) SetParentPath(parentPath string) {
	m.parentPath = parentPath
}

func (m *Metadata) SetPath(path string) {
	m.path = path
}

func (m *Metadata) SetType(mtype MetadataType) {
	m.mtype = mtype
}

func (m *Metadata) SetINodeID(iNodeID uuid.UUID) {
	m.iNodeID = iNodeID
}

func (m *Metadata) GetHash() string {
	return m.hash
}

func (m *Metadata) SetHash(val string) {
	m.hash = val
}
