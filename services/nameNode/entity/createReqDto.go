package entity

import "path/filepath"

type CreateReqDto struct {
	replicationTarget uint32
	blockSplitTarget  uint32
	fileSize          uint64
	leaseTimeInSec    uint64
	path              string
}

// Getter methods
func (c *CreateReqDto) GetReplicationTarget() uint32 {
	return c.replicationTarget
}

func (c *CreateReqDto) GetBlockSplitTarget() uint32 {
	return c.blockSplitTarget
}

func (c *CreateReqDto) GetFileSize() uint64 {
	return c.fileSize
}

func (c *CreateReqDto) GetLeaseTimeInSec() uint64 {
	return c.leaseTimeInSec
}

func (c *CreateReqDto) GetPath() string {
	return c.path
}

func (c *CreateReqDto) GetParentPath() string {
	return filepath.Dir(c.path)
}

// Setter methods
func (c *CreateReqDto) SetReplicationTarget(replicationTarget uint32) {
	c.replicationTarget = replicationTarget
}

func (c *CreateReqDto) SetBlockSplitTarget(blockSplitTarget uint32) {
	c.blockSplitTarget = blockSplitTarget
}

func (c *CreateReqDto) SetFileSize(fileSize uint64) {
	c.fileSize = fileSize
}

func (c *CreateReqDto) SetLeaseTimeInSec(leaseTimeInSec uint64) {
	c.leaseTimeInSec = leaseTimeInSec
}

func (c *CreateReqDto) SetPath(path string) {
	c.path = path
}
