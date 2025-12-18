package main

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"
)

type Cache struct {
	podID string

	// L1
	l1 map[string]string
	mu sync.RWMutex

	// L2
	rdb *redis.Client

	// Deduplicação por pod
	sf singleflight.Group
}

func NewCache(addr, podID string) *Cache {
	logPod(podID, "cache", "initializing cache")

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	return &Cache{
		podID: podID,
		l1:    make(map[string]string),
		rdb:   rdb,
	}
}

func (c *Cache) Get(
	ctx context.Context,
	key string,
	ttl time.Duration,
	loader func() string,
) string {

	// 1️⃣ L1 cache
	c.mu.RLock()
	val, ok := c.l1[key]
	c.mu.RUnlock()

	if ok {
		logPod(c.podID, "cache", "L1 HIT key="+key)
		return val
	}

	// 2️⃣ Deduplicação por pod
	v, _, _ := c.sf.Do(key, func() (interface{}, error) {

		// Double check L1
		c.mu.RLock()
		val, ok := c.l1[key]
		c.mu.RUnlock()
		if ok {
			return val, nil
		}

		// 3️⃣ L2 Redis
		val, err := c.rdb.Get(ctx, key).Result()
		if err == nil {
			logPod(c.podID, "cache", "L2 HIT key="+key)

			c.mu.Lock()
			c.l1[key] = val
			c.mu.Unlock()

			return val, nil
		}

		// 4️⃣ Loader
		logPod(c.podID, "loader", "loading key="+key)
		value := loader()

		c.rdb.Set(ctx, key, value, ttl)

		c.mu.Lock()
		c.l1[key] = value
		c.mu.Unlock()

		logPod(c.podID, "cache", "SET L1+L2 key="+key)

		return value, nil
	})

	return v.(string)
}
