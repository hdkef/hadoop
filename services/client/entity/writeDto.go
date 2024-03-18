package entity

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"strconv"

	clientProto "github.com/hdkef/hadoop/proto/client"
)

type CreateDto struct {
	replicationTarget uint32
	blockSplitTarget  uint32
	file              []byte
	leaseTimeInSec    uint32
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

func (c *CreateDto) SetLeaseTimeInSec(value uint32) {
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

func (c *CreateDto) GetLeaseTimeInSec() uint32 {
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

func (w *CreateDto) NewFromHttp(r *http.Request) error {

	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum
	if err != nil {
		return err
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	w.file = data
	w.path = r.FormValue("path")
	blockSplitTarget := r.FormValue("blockSplitTarget")
	blockSplitTargetVal := 0

	if blockSplitTarget != "" {
		blockSplitTargetVal, err = strconv.Atoi(blockSplitTarget)
		if err != nil {
			return err
		}
	}

	replicationTarget := r.FormValue("replicationTarget")
	replicationTargetVal := 0
	if replicationTarget != "" {
		replicationTargetVal, err = strconv.Atoi(replicationTarget)
		if err != nil {
			return err
		}
	}

	w.replicationTarget = uint32(replicationTargetVal)
	w.blockSplitTarget = uint32(blockSplitTargetVal)

	return nil
}
