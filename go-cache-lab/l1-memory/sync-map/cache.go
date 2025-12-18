package main

import "sync"

// Cache implementa um cache L1 em memória
// usando sync.Map, que já é thread-safe.
type Cache struct {
	data sync.Map
}

func NewCache() *Cache {
	logger.Println("[cache] initializing cache")
	return &Cache{}
}

// Get não precisa de lock explícito
func (c *Cache) Get(key string) (string, bool) {
	value, ok := c.data.Load(key)
	if ok {
		logger.Printf("[cache] HIT key=%s\n", key)
		return value.(string), true
	}

	logger.Printf("[cache] MISS key=%s\n", key)
	return "", false
}

// Set também é thread-safe
func (c *Cache) Set(key, value string) {
	logger.Printf("[cache] SET key=%s\n", key)
	c.data.Store(key, value)
}
