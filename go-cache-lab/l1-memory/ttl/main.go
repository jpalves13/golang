package main

import (
	"time"
)

func main() {
	logger.Println("[app] starting application")

	cache := NewCache()

	cache.Set("key-1", "value", 2*time.Second)

	// HIT
	cache.Get("key-1")

	time.Sleep(3 * time.Second)

	// EXPIRED
	cache.Get("key-1")

	logger.Println("[app] finished")
}
