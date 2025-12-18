package main

import (
	"fmt"
	"sync"
)

func main() {
	logger.Println("[app] starting application")

	cache := NewCache()
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			key := fmt.Sprintf("key-%d", i%3)

			cache.Set(key, "value")
			cache.Get(key)
			cache.Get(key)
		}(i)
	}

	wg.Wait()
	logger.Println("[app] finished")
}
