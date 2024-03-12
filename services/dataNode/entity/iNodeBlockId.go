package entity

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

type INodeBlockID struct {
	iNodeID uuid.UUID
	blockID uuid.UUID
}

// Set methods allow setting individual fields of INodeBlockID
func (ib *INodeBlockID) SetINodeID(iNodeID uuid.UUID) {
	ib.iNodeID = iNodeID
}

func (ib *INodeBlockID) SetBlockID(blockID uuid.UUID) {
	ib.blockID = blockID
}

// Get methods allow getting individual fields of INodeBlockID
func (ib *INodeBlockID) GetINodeID() uuid.UUID {
	return ib.iNodeID
}

func (ib *INodeBlockID) GetBlockID() uuid.UUID {
	return ib.blockID
}

func (i *INodeBlockID) GetKey() string {
	return fmt.Sprintf("inodeblockid_%s_%s", i.iNodeID, i.blockID)
}

func (i *INodeBlockID) Write(root string, binaryData []byte) error {

	dirPath := fmt.Sprintf("%s/%s.bin", root, i.iNodeID.String()+i.blockID.String())

	err := os.WriteFile(dirPath, binaryData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (i *INodeBlockID) Remove(root string) error {

	dirPath := fmt.Sprintf("%s/%s.bin", root, i.iNodeID.String()+i.blockID.String())

	err := os.Remove(dirPath)
	if err != nil {
		return err
	}

	return nil
}
