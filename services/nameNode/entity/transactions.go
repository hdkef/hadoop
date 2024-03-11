package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	TRANSACTION_ACTION_CREATE TransactionsAction = 1
	TRANSACTION_ACTION_UPDATE TransactionsAction = 2
	TRANSACTION_ACTION_DELETE TransactionsAction = 3
)

type TransactionsAction int

type Transactions struct {
	id                  uuid.UUID
	action              TransactionsAction
	createdAt           time.Time
	isCommitted         bool
	transactionNodeInfo []*TransactionNodeInfo
	leaseTimeInSecond   uint64
	metadata            *Metadata
}

func (t *Transactions) GetID() uuid.UUID {
	return t.id
}

func (t *Transactions) GetAction() TransactionsAction {
	return t.action
}

func (t *Transactions) GetCreatedAt() time.Time {
	return t.createdAt
}

func (t *Transactions) GetIsCommitted() bool {
	return t.isCommitted
}

func (t *Transactions) GetTransactionNodeInfo() []*TransactionNodeInfo {
	return t.transactionNodeInfo
}

func (t *Transactions) GetLeaseTimeInSecond() uint64 {
	return t.leaseTimeInSecond
}

// Setter methods
func (t *Transactions) SetID(id uuid.UUID) {
	t.id = id
}

func (t *Transactions) SetAction(action TransactionsAction) {
	t.action = action
}

func (t *Transactions) SetCreatedAt(createdAt time.Time) {
	t.createdAt = createdAt
}

func (t *Transactions) SetIsCommitted(isCommitted bool) {
	t.isCommitted = isCommitted
}

func (t *Transactions) SetTransactionNodeInfo(nodeInfo []*TransactionNodeInfo) {
	t.transactionNodeInfo = nodeInfo
}

func (t *Transactions) SetLeaseTimeInSecond(leaseTime uint64) {
	t.leaseTimeInSecond = leaseTime
}

func (t *Transactions) SetMetadata(metadata *Metadata) {
	t.metadata = metadata
}

func (t *Transactions) GetMetadata() *Metadata {
	return t.metadata
}
