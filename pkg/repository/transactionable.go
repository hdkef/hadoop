package repository

type Transactionable interface {
	Commit() error
	Rollback() error
}
