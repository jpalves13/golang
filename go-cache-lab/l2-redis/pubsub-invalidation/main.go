package main

import (
	"context"
	"sync"
	"time"
)

func main() {
	logApp("starting pubsub invalidation simulation")

	ctx := context.Background()
	redisAddr := "localhost:6379"

	pods := []*Cache{
		NewCache(redisAddr, "pod-1"),
		NewCache(redisAddr, "pod-2"),
		NewCache(redisAddr, "pod-3"),
	}

	key := "user:42"

	// Pod 1 escreve
	pods[0].Set(ctx, key, "john-doe", 30*time.Second)

	time.Sleep(300 * time.Millisecond)

	// Todos leem
	var wg sync.WaitGroup
	for _, pod := range pods {
		wg.Add(1)

		go func(p *Cache) {
			defer wg.Done()
			p.Get(ctx, key)
		}(pod)
	}

	wg.Wait()

	logApp("finished simulation")
}
