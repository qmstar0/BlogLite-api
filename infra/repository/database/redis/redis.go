package redis

import (
	"blog/infra/config"
	"blog/infra/repository/dao"
	"context"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

var (
	CC *CacheClient
	_  dao.CacheClient = &CacheClient{}
)

// CacheClient 缓存器
type CacheClient struct {
	client *redis.Client
}

// GetCacheClient 获取缓存器
func GetCacheClient() *CacheClient {
	return CC
}

// 模块初始化
func init() {
	pwd := os.Getenv("BLOG_REDIS_PASSWORD")
	if pwd == "" {
		panic("redis is not configured: see env:BLOG_REDIS_PASSWORD")
	}
	CC = &CacheClient{
		redis.NewClient(&redis.Options{
			Addr: config.Conf.Redis.Addr, // Redis 服务器的地址和端口
			//Password: pwd,
		}),
	}
	_, err := CC.client.Ping(context.Background()).Result()
	if err != nil {
		panic("无法连接到 Redis:" + err.Error())
	}
}

// DelCache 删除
func (c *CacheClient) DelCache(ctx context.Context, keys ...string) error {
	return c.client.Del(ctx, keys...).Err()
}

// SetCache 设置
func (c *CacheClient) SetCache(ctx context.Context, key string, modelObj any, expiration time.Duration) error {
	return c.client.Set(ctx, key, modelObj, expiration).Err()
}

// GetCache 获取
func (c *CacheClient) GetCache(ctx context.Context, key string, modelObj any) error {
	return c.client.Get(ctx, key).Scan(modelObj)
}

// Close 关闭连接
func Close() {
	_ = CC.client.Close()
}
