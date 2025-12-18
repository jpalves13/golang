package main

import (
	"fmt"
	"sync"
)

func main() {
	logger.Println("[app] starting application")

	cache := NewCache()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			key := fmt.Sprintf("key-%d", i%3)

			// Simula workload read-heavy
			cache.Set(key, "value")
			cache.Get(key)
			cache.Get(key)
			cache.Get(key)
		}(i)
	}

	wg.Wait()
	logger.Println("[app] finished")
}
