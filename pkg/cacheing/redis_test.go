package cacheing_test

import (
	"blog/pkg/cacheing"
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
	"testing"
	"time"
)

func TestNewRedisCache(t *testing.T) {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.10:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	var (
		wg      = sync.WaitGroup{}
		ctx     = context.Background()
		cache   = cacheing.NewRedisCache(redisCli)
		mapData = map[string]string{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
		}
	)

	_, err := redisCli.Ping(ctx).Result()
	if err != nil {
		t.Log("content success!")
	}

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
