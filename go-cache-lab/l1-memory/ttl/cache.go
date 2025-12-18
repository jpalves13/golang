package main

import (
	"sync"
	"time"
)

// item representa um valor armazenado no cache
// junto com seu tempo de expiração.
type item struct {
	value     string
	expiresAt time.Time
}

type Cache struct {
	mu   sync.RWMutex
	data map[string]item
}

func NewCache() *Cache {
	logger.Println("[cache] initializing cache with TTL")
	return &Cache{
		data: make(map[string]item),
	}
}

// Set armazena um valor com TTL
func (c *Cache) Set(key, value string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = item{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}

	logger.Printf("[cache] SET key=%s ttl=%s\n", key, ttl)
}

// Get retorna o valor apenas se ainda estiver válido
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	it, ok := c.data[key]
	c.mu.RUnlock()

	if !ok {
		logger.Printf("[cache] MISS key=%s\n", key)
		return "", false
	}

	// Verifica expiração
	if time.Now().After(it.expiresAt) {
		logger.Printf("[cache] EXPIRED key=%s\n", key)

		// Remove item expirado
		c.mu.Lock()
		delete(c.data, key)
		c.mu.Unlock()

		return "", false
	}

	logger.Printf("[cache] HIT key=%s\n", key)
	return it.value, true
}
