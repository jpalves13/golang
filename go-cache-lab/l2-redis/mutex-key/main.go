package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	log.Println("[app] starting pod simulation")

	ctx := context.Background()
	var wg sync.WaitGroup

	pods := 4
	key := "hot-key"

	for i := 1; i <= pods; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			podID := fmt.Sprintf("pod-%d", id)
			cache := NewCache("localhost:6379", podID)

			cache.GetOrSet(ctx, key, 10*time.Second, func() string {
				log.Printf("[%s][loader] generating value\n", podID)
				time.Sleep(1 * time.Second)
				return "value-from-" + podID
			})
		}(i)
	}

	wg.Wait()
	log.Println("[app] finished simulation")
}
