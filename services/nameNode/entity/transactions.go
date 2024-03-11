package entity

import "time"

const (
	TRANSACTION_ACTION_CREATE TransactionsAction = 1
	TRANSACTION_ACTION_UPDATE TransactionsAction = 2
	TRANSACTION_ACTION_DELETE TransactionsAction = 3
)

type TransactionsAction int

type Transactions struct {
	id                  string
	action              TransactionsAction
	createdAt           time.Time
	isCommitted         bool
	transactionNodeInfo []*TransactionNodeInfo
	leaseTimeInSecond   int64
}
