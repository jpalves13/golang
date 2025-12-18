package main

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisLock struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisLock(client *redis.Client, ttl time.Duration) *RedisLock {
	return &RedisLock{
		client: client,
		ttl:    ttl,
	}
}

func (l *RedisLock) Lock(ctx context.Context, key string) bool {
	lockKey := "lock:" + key
	ok, _ := l.client.SetNX(ctx, lockKey, "1", l.ttl).Result()
	return ok
}

func (l *RedisLock) Unlock(ctx context.Context, key string) {
	lockKey := "lock:" + key
	l.client.Del(ctx, lockKey)
}
