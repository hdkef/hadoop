package entity

type CreateReqDto struct {
	ReplicationTarget uint32
	BlockSplitTarget  uint32
	FileSize          uint64
}
