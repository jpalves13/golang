package main

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	podID string

	l1     map[string]string
	mu     sync.RWMutex
	redis  *redis.Client
	pubsub *redis.PubSub
}

func NewCache(addr, podID string) *Cache {
	logPod(podID, "cache", "initializing cache")

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		logPod(podID, "cache", "redis not available")
		panic(err)
	}

	c := &Cache{
		podID: podID,
		l1:    make(map[string]string),
		redis: rdb,
	}

	// Subscribe to invalidation channel
	c.pubsub = rdb.Subscribe(context.Background(), "cache-invalidate")

	go c.listenInvalidation()

	return c
}

// Listener de invalidação
func (c *Cache) listenInvalidation() {
	ch := c.pubsub.Channel()

	for msg := range ch {
		key := msg.Payload

		c.mu.Lock()
		delete(c.l1, key)
		c.mu.Unlock()

		logPod(c.podID, "invalidation", "L1 invalidated key="+key)
	}
}

// Get com L1 + L2
func (c *Cache) Get(ctx context.Context, key string) string {
	// L1
	c.mu.RLock()
	val, ok := c.l1[key]
	c.mu.RUnlock()

	if ok {
		logPod(c.podID, "cache", "L1 HIT key="+key)
		return val
	}

	// L2
	val, err := c.redis.Get(ctx, key).Result()
	if err == nil {
		logPod(c.podID, "cache", "L2 HIT key="+key)

		c.mu.Lock()
		c.l1[key] = val
		c.mu.Unlock()

		return val
	}

	logPod(c.podID, "cache", "MISS key="+key)
	return ""
}

// Set + publish invalidation
func (c *Cache) Set(ctx context.Context, key, value string, ttl time.Duration) {
	c.redis.Set(ctx, key, value, ttl)

	// Publish invalidation
	c.redis.Publish(ctx, "cache-invalidate", key)

	logPod(c.podID, "cache", "SET key="+key+" (invalidate others)")
}
