package repository

import (
	"context"
	"time"
)

type KeyValueRepository interface {
	Set(ctx context.Context, key string, value []byte, exp *time.Duration) error
	Get(ctx context.Context, key string) ([]byte, error)
	Del(ctx context.Context, key string) error
	Incr(ctx context.Context, key string) error
	Decr(ctx context.Context, key string) error
	CloseConn()
}
