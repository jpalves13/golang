package main

import "fmt"

func worker(id int, cache *Cache) {
	key := fmt.Sprintf("key-%d", id%5)
	cache.Set(key, "value")
	cache.Get(key)
}
