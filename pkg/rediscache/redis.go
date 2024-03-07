package rediscache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

type redisCache struct {
	cli *redis.Client
}

var cache Cache

func GetCacher() Cache {
	if cache == nil {
		panic("cache(redis) not init")
	}
	return cache
}

func Init() (closeFn func() error) {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("CACHE_HOST_DEV"),
		Password: os.Getenv("CACHE_PASSWORD_DEV"), // 没有密码，默认值
		DB:       0,                               // 默认DB 0
	})
	_, err := redisCli.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
	cache = redisCache{cli: redisCli}

	return redisCli.Close
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
