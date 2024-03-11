package entity

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hdkef/hadoop/pkg/helper"
)

type INodeBlockID struct {
	iNodeID string
	blockID string
	dirPath string
}

// Set methods allow setting individual fields of INodeBlockID
func (ib *INodeBlockID) SetINodeID(iNodeID string) {
	ib.iNodeID = iNodeID
}

func (ib *INodeBlockID) SetBlockID(blockID string) {
	ib.blockID = blockID
}

func (ib *INodeBlockID) SetDirPath(dirPath string) {
	ib.dirPath = dirPath
}

// Get methods allow getting individual fields of INodeBlockID
func (ib *INodeBlockID) GetINodeID() string {
	return ib.iNodeID
}

func (ib *INodeBlockID) GetBlockID() string {
	return ib.blockID
}

func (ib *INodeBlockID) GetDirPath() string {
	return ib.dirPath
}

func (i *INodeBlockID) GetKey() string {
	return fmt.Sprintf("inodeblockid_%s_%s", i.iNodeID, i.blockID)
}

func (i *INodeBlockID) Write(root string, binaryData []byte) error {

	randomStr, err := helper.GenerateRandomString()
	if err != nil {
		return err
	}

	i.dirPath = fmt.Sprintf("%s/%s.bin", root, randomStr)

	err = os.WriteFile(i.dirPath, binaryData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (i *INodeBlockID) ToJSON() []byte {

	b, _ := json.Marshal(i)

	return b
}
