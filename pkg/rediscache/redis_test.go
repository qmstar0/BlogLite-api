package rediscache_test

import (
	"blog/pkg/rediscache"
	"context"
	"github.com/go-redis/redis/v8"
	"os"
	"sync"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("CACHE_HOST_DEV"),
		Password: os.Getenv("CACHE_PASSWORD_DEV"), // 没有密码，默认值
		DB:       0,                               // 默认DB 0
	})
	_, err := redisCli.Ping(context.TODO()).Result()
	if err != nil {
		t.Fatal("content err:", err)
	}
}

func TestNewRedisCache(t *testing.T) {

	closeFn := rediscache.Init()
	defer closeFn()

	var (
		err     error
		wg      = sync.WaitGroup{}
		ctx     = context.Background()
		cache   = rediscache.GetCacher()
		mapData = map[string]string{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
		}
	)

	go func() {
		wg.Add(1)
		err = cache.Set(ctx, "k_4s", "v", time.Second*1)
		if err != nil {
			t.Error("err on set:", "k_4s", err)
		}
		time.AfterFunc(time.Second*2, func() {
			v, err := cache.Get(ctx, "k_4s")
			if err == nil {
				t.Error("no err while getting timeouted data", v)
			}
			wg.Done()
		})
		v, err := cache.Get(ctx, "k_4s")
		if err != nil {
			t.Error("err on get:", "k_4s", err)
		}
		if v != "v" {
			t.Error("data err:", v, "!= v")
		}
	}()

	err = cache.SetBytes(ctx, "k_b", []byte("hi"), time.Second*2)
	if err != nil {
		t.Error("err on set bytes", err)
	}
	bytes, err := cache.GetBytes(ctx, "k_b")
	if err != nil {
		t.Error("err on get bytes", err)
	}
	if string(bytes) != "hi" {
		t.Error(bytes, " != hi")
	}

	err = cache.SetMap(ctx, "k_m", mapData, time.Second*2)
	if err != nil {
		t.Error("err on set map", err)
	}
	data, err := cache.GetMap(ctx, "k_m")
	if err != nil {
		t.Error("err on get map", err)
	}
	for k, v := range mapData {
		if data[k] != v {
			t.Error("err on map data k != v")
		}
	}
	err = cache.SetMapKV(ctx, "k_m", "k4", "v4")
	if err != nil {
		t.Error("err on set map kv", err)
	}
	v, err := cache.GetMapKV(ctx, "k_m", "k4")
	if err != nil {
		t.Error("err on get map kv", err)
	}
	if v != "v4" {
		t.Error("geted data err")
	}
	getMap, err := cache.GetMap(ctx, "k_m")
	if err != nil {
		t.Error("err on get map", err)
	}
	if len(getMap) != 4 {
		t.Error("geted data length err")
	}
	wg.Wait()
}
