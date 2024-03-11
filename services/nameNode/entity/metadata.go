package entity

const (
	METADATA_TYPE_DIR  = 1
	METADATA_TYPE_FILE = 2
)

type MetadataType uint8

type Metadata struct {
	parentPath string
	path       string
	mtype      MetadataType
	iNodeId    string
}
