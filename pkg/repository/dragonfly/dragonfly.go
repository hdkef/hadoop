package dragonfly

import (
	"context"
	"errors"
	"time"

	"github.com/hdkef/hadoop/pkg/repository"
	"github.com/redis/go-redis/v9"
)

type DragonFlyConfig struct {
	Addr     string
	Password string
	DB       int
}

type DragonFlyRepo struct {
	db *redis.Client
}

// CloseConn implements repository.KeyValueRepository.
func (d *DragonFlyRepo) CloseConn() {
	d.db.Close()
}

// Decr implements repository.KeyValueRepository.
func (d *DragonFlyRepo) Decr(ctx context.Context, key string, amount int) error {
	return d.db.DecrBy(ctx, key, int64(amount)).Err()
}

// Del implements repository.KeyValueRepository.
func (d *DragonFlyRepo) Del(ctx context.Context, key string) error {
	return d.db.Del(ctx, key).Err()
}

// Get implements repository.KeyValueRepository.
func (d *DragonFlyRepo) Get(ctx context.Context, key string) ([]byte, error) {
	return d.db.Get(ctx, key).Bytes()
}

// Incr implements repository.KeyValueRepository.
func (d *DragonFlyRepo) Incr(ctx context.Context, key string, amount int) error {
	return d.db.IncrBy(ctx, key, int64(amount)).Err()
}

// Set implements repository.KeyValueRepository.
func (d *DragonFlyRepo) Set(ctx context.Context, key string, value []byte, exp *time.Duration) error {

	withTTL := time.Duration(0)
	if exp != nil {
		withTTL = *exp
	}

	return d.db.Set(ctx, key, value, withTTL).Err()
}

func NewDragonFlyRepo(cfg *DragonFlyConfig) repository.KeyValueRepository {

	db := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if db == nil {
		panic(errors.New("db err"))
	}

	return &DragonFlyRepo{
		db: db,
	}
}
