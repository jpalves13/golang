package main

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func NewCache(addr string) *Cache {
	log.Println("[cache] initializing redis cache")

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("[cache] redis not available: %v", err)
	}

	return &Cache{client: rdb}
}

func (c *Cache) Get(ctx context.Context, key string) (string, bool) {
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Printf("[cache] MISS key=%s\n", key)
		return "", false
	}
	if err != nil {
		log.Printf("[cache] ERROR key=%s err=%v\n", key, err)
		return "", false
	}

	log.Printf("[cache] HIT key=%s\n", key)
	return val, true
}

func (c *Cache) Set(ctx context.Context, key, value string, ttl time.Duration) {
	log.Printf("[cache] SET key=%s\n", key)
	c.client.Set(ctx, key, value, ttl)
}
