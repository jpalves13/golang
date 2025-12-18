package main

import (
	"context"
	"sync"
	"time"
)

func main() {
	logApp("starting l1-l2-singleflight simulation")

	ctx := context.Background()
	redisAddr := "localhost:6379"
	key := "hot-key"

	pods := []*Cache{
		NewCache(redisAddr, "pod-1"),
		NewCache(redisAddr, "pod-2"),
	}

	var wg sync.WaitGroup

	for _, pod := range pods {
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func(c *Cache) {
				defer wg.Done()

				c.Get(ctx, key, 15*time.Second, func() string {
					time.Sleep(500 * time.Millisecond)
					return "expensive-value"
				})
			}(pod)
		}
	}

	wg.Wait()
	logApp("finished simulation")
}
