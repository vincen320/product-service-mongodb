package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func FindKey(ctx context.Context, rdb *redis.Client, key string) (string, error) {
	var result *redis.StringCmd
	_, err := rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		result = pipe.Get(ctx, key)
		return nil
	})

	if err != nil {
		return result.Val(), err
	}
	return result.Val(), nil
}

func SetProductCache(ctx context.Context, rdb *redis.Client, key string, value interface{}) error {
	var status *redis.StatusCmd
	_, err := rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		ttl := time.Second * 30
		status = pipe.Set(ctx, key, value, ttl)
		return nil
	})

	if err != nil {
		return err
	}
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}
