package main

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
	lock   *RedisLock
	podID  string
}

func NewCache(addr, podID string) *Cache {
	log.Printf("[%s][cache] initializing redis cache\n", podID)

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("[%s][cache] redis not available: %v", podID, err)
	}

	return &Cache{
		client: rdb,
		lock:   NewRedisLock(rdb, 5*time.Second),
		podID:  podID,
	}
}

func (c *Cache) GetOrSet(
	ctx context.Context,
	key string,
	ttl time.Duration,
	loader func() string,
) string {

	// 1️⃣ Tentativa de cache
	val, err := c.client.Get(ctx, key).Result()
	if err == nil {
		log.Printf("[%s][cache] HIT key=%s\n", c.podID, key)
		return val
	}

	log.Printf("[%s][cache] MISS key=%s\n", c.podID, key)

	// 2️⃣ Lock distribuído
	if c.lock.Lock(ctx, key) {
		log.Printf("[%s][lock] ACQUIRED key=%s\n", c.podID, key)
		defer c.lock.Unlock(ctx, key)

		// Double check
		if val, err := c.client.Get(ctx, key).Result(); err == nil {
			return val
		}

		value := loader()
		c.client.Set(ctx, key, value, ttl)

		log.Printf("[%s][cache] SET key=%s\n", c.podID, key)
		return value
	}

	// 3️⃣ Espera outro pod preencher
	log.Printf("[%s][lock] WAIT key=%s\n", c.podID, key)
	time.Sleep(200 * time.Millisecond)

	val, _ = c.client.Get(ctx, key).Result()
	return val
}
