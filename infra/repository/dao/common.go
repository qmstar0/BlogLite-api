package dao

import (
	"blog/infra/repository/database/mysql"
	"context"
	"time"
)

// CacheClient 缓存需要实现的接口
type CacheClient interface {
	DelCache(ctx context.Context, keys ...string) error
	SetCache(ctx context.Context, key string, modelObj any, expiration time.Duration) error
	GetCache(ctx context.Context, key string, modelObj any) error
}

// Dao 用户仓储dao层
type Dao struct {
	db    *mysql.DBClient
	cache CacheClient
}

// NewDao 构建
func NewDao(db *mysql.DBClient, cache CacheClient) *Dao {
	return &Dao{db: db, cache: cache}
}
