package entity

import "fmt"

type NodeStorage struct {
	nodeID                  string
	leaseUsedStorage        uint64
	actualUsedStorage       uint64
	totalStorage            uint64
	allocatedHelper         bool
	allocatedAmountHelper   uint32
	leaseUsedStorageChanged bool
}

// Setter methods
func (n *NodeStorage) SetNodeID(nodeID string) {
	n.nodeID = nodeID
}

func (n *NodeStorage) SetLeaseUsedStorage(leaseUsedStorage uint64) {
	n.leaseUsedStorage = leaseUsedStorage
}

func (n *NodeStorage) SetActualUsedStorage(actualUsedStorage uint64) {
	n.actualUsedStorage = actualUsedStorage
}

// Getter methods
func (n *NodeStorage) GetNodeID() string {
	return n.nodeID
}

func (n *NodeStorage) GetLeaseUsedStorage() uint64 {
	return n.leaseUsedStorage
}

func (n *NodeStorage) GetActualUsedStorage() uint64 {
	return n.actualUsedStorage
}

func (n *NodeStorage) IsAllocated() bool {
	return n.allocatedHelper
}

func (n *NodeStorage) SetAllocated() {
	n.allocatedHelper = true
	n.allocatedAmountHelper += 1
}

func (n *NodeStorage) DecrementLeaseStorage(val uint64) {

	if val <= n.leaseUsedStorage {
		n.leaseUsedStorage -= val
	}
}

func (n *NodeStorage) IncrementLeaseStorage(val uint64) {

	n.leaseUsedStorage += val
}

func (n *NodeStorage) GetAllocated() uint32 {
	return n.allocatedAmountHelper
}

func (n *NodeStorage) SetTotalStorage(f uint64) {
	n.totalStorage = f
}

func (n *NodeStorage) SetLeasedUsedStorageChanged() {
	n.leaseUsedStorageChanged = true
}

func (n *NodeStorage) GetLeasedUsedStorageChanged() bool {
	return n.leaseUsedStorageChanged
}

func (n *NodeStorage) GetTotalStorage() uint64 {
	return n.totalStorage
}

func (n *NodeStorage) GetFreeSpace() uint64 {
	return n.totalStorage - (n.actualUsedStorage + n.leaseUsedStorage)
}

func (n *NodeStorage) GenerateKey() string {
	return fmt.Sprintf("NodeStorage_%s", n.nodeID)
}
