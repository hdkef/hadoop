package entity

import (
	"crypto/sha256"
	"encoding/hex"

	clientProto "github.com/hdkef/hadoop/proto/client"
)

type CreateDto struct {
	replicationTarget uint32
	blockSplitTarget  uint32
	file              []byte
	leaseTimeInSec    uint64
	path              string
}

func (w *CreateDto) Tokenize(numParts int) [][]byte {
	// Calculate the size of each part
	partSize := (len(w.file) + numParts - 1) / numParts // Round up to ensure coverage of all bytes
	tokens := make([][]byte, 0)
	for i := 0; i < len(w.file); i += partSize {
		end := i + partSize
		if end > len(w.file) {
			end = len(w.file)
		}
		tokens = append(tokens, w.file[i:end])
	}
	return tokens
}

func (w *CreateDto) NewFromProto(req *clientProto.CreateReq) error {
	w.file = req.GetFiles()
	w.path = req.GetPath()
	w.blockSplitTarget = req.GetBlockSplitTarget()
	w.replicationTarget = req.GetReplicationTarget()

	return nil
}

// Setter methods
func (c *CreateDto) SetReplicationTarget(value uint32) {
	c.replicationTarget = value
}

func (c *CreateDto) SetBlockSplitTarget(value uint32) {
	c.blockSplitTarget = value
}

func (c *CreateDto) SetFile(value []byte) {
	c.file = value
}

func (c *CreateDto) SetLeaseTimeInSec(value uint64) {
	c.leaseTimeInSec = value
}

func (c *CreateDto) SetPath(value string) {
	c.path = value
}

// Getter methods
func (c *CreateDto) GetReplicationTarget() uint32 {
	return c.replicationTarget
}

func (c *CreateDto) GetBlockSplitTarget() uint32 {
	return c.blockSplitTarget
}

func (c *CreateDto) GetFile() []byte {
	return c.file
}

func (c *CreateDto) GetLeaseTimeInSec() uint64 {
	return c.leaseTimeInSec
}

func (c *CreateDto) GetPath() string {
	return c.path
}

func (c *CreateDto) GetHashFile() string {
	hash := sha256.Sum256(c.file)

	// Convert checksum to hexadecimal string
	return hex.EncodeToString(hash[:])
}

func (c *CreateDto) GetFileSize() uint64 {
	return uint64(len(c.file))
}
