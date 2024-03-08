package dragonfly

import (
	"context"
	"fmt"
	"time"

	badger "github.com/dgraph-io/badger/v4"
	"github.com/hdkef/hadoop/pkg/repository"
)

type BadgerRepo struct {
	db *badger.DB
}

// CloseConn implements repository.KeyValueRepository.
func (b *BadgerRepo) CloseConn() {
	b.db.Close()
}

func NewBadgerRepo(storageRoot string) repository.KeyValueRepository {
	db, err := badger.Open(badger.DefaultOptions(fmt.Sprintf("%s/badger", storageRoot)))
	if err != nil {
		panic(err)
	}
	return &BadgerRepo{
		db: db,
	}
}

// Decr implements repository.KeyValueRepository.
func (b *BadgerRepo) Decr(ctx context.Context, key string) error {
	panic("unimplemented")
}

// Del implements repository.KeyValueRepository.
func (b *BadgerRepo) Del(ctx context.Context, key string) error {
	panic("unimplemented")
}

// Get implements repository.KeyValueRepository.
func (b *BadgerRepo) Get(ctx context.Context, key string) (data []byte, err error) {
	err = b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		// Retrieve value as []byte
		data, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// Incr implements repository.KeyValueRepository.
func (b *BadgerRepo) Incr(ctx context.Context, key string) error {
	panic("unimplemented")
}

// Set implements repository.KeyValueRepository.
func (b *BadgerRepo) Set(ctx context.Context, key string, value []byte, exp *time.Time) error {
	return b.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), value)
		return err
	})
}
