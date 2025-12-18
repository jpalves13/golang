package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	log.Println("[app] starting application")

	ctx := context.Background()

	cache := NewCache("localhost:6379")

	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key-%d", i%2)

		if _, ok := cache.Get(ctx, key); !ok {
			cache.Set(ctx, key, "value", 10*time.Second)
		}
	}

	log.Println("[app] finished")
}
