package entity

import (
	"path/filepath"

	nameNodeProto "github.com/hdkef/hadoop/proto/nameNode"
)

type CreateReqDto struct {
	replicationTarget uint32
	blockSplitTarget  uint32
	fileSize          uint64
	leaseTimeInSec    uint32
	path              string
	hash              string
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

func (c *CreateReqDto) GetLeaseTimeInSec() uint32 {
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

func (c *CreateReqDto) SetLeaseTimeInSec(leaseTimeInSec uint32) {
	c.leaseTimeInSec = leaseTimeInSec
}

func (c *CreateReqDto) SetPath(path string) {
	c.path = path
}

func (c *CreateReqDto) GetHash() string {
	return c.hash
}

func (c *CreateReqDto) SetHash(val string) {
	c.hash = val
}

func (c *CreateReqDto) FromProto(req *nameNodeProto.QueryNodeTargetCreateReq) error {

	c.blockSplitTarget = req.GetBlockSplitTarget()
	c.fileSize = req.GetFileSize()
	c.replicationTarget = req.GetReplicationTarget()
	c.hash = req.GetHash()
	c.path = req.GetPath()
	c.leaseTimeInSec = req.GetLeaseTimeInSec()

	return nil
}
