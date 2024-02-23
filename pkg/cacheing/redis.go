package cacheing

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type redisCache struct {
	cli *redis.Client
}

func NewRedisCache(client *redis.Client) Cache {
	return &redisCache{cli: client}
}

func (c redisCache) Set(ctx context.Context, k, v string, timeout time.Duration) error {
	return c.cli.Set(ctx, k, v, timeout).Err()
}

func (c redisCache) Get(ctx context.Context, k string) (string, error) {
	return c.cli.Get(ctx, k).Result()
}

func (c redisCache) SetMap(ctx context.Context, mapName string, mapEntity map[string]string, timeout time.Duration) error {
	err := c.cli.HMSet(ctx, mapName, mapEntity).Err()
	if err != nil {
		return err
	}
	if timeout > 0 {
		return c.cli.Expire(ctx, mapName, timeout).Err()
	}
	return nil
}

func (c redisCache) GetMap(ctx context.Context, mapName string) (map[string]string, error) {
	return c.cli.HGetAll(ctx, mapName).Result()
}

func (c redisCache) SetBytes(ctx context.Context, k string, v []byte, timeout time.Duration) error {
	return c.cli.Set(ctx, k, v, timeout).Err()
}

func (c redisCache) GetBytes(ctx context.Context, k string) ([]byte, error) {
	return c.cli.Get(ctx, k).Bytes()
}

func (c redisCache) SetMapKV(ctx context.Context, mapName, k, v string) error {
	return c.cli.HSet(ctx, mapName, k, v).Err()
}

func (c redisCache) GetMapKV(ctx context.Context, mapName, k string) (string, error) {
	return c.cli.HGet(ctx, mapName, k).Result()
}
