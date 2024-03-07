package rediscache

import (
	"context"
	"time"
)

type Cache interface {
	Set(ctx context.Context, k, v string, timeout time.Duration) error
	Get(ctx context.Context, k string) (string, error)
	SetMap(ctx context.Context, mapName string, mapEntity map[string]string, timeout time.Duration) error
	GetMap(ctx context.Context, mapName string) (map[string]string, error)
	SetBytes(ctx context.Context, k string, v []byte, timeout time.Duration) error
	GetBytes(ctx context.Context, k string) ([]byte, error)
	SetMapKV(ctx context.Context, mapName, k, v string) error
	GetMapKV(ctx context.Context, mapName, k string) (string, error)
}
