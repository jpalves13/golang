package main

import "sync"

// Cache implementa um cache L1 em memória
// usando map protegido por sync.RWMutex.
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewCache() *Cache {
	logger.Println("[cache] initializing cache")
	return &Cache{
		data: make(map[string]string),
	}
}

// Get permite múltiplas leituras concorrentes
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.data[key]
	if ok {
		logger.Printf("[cache] HIT key=%s\n", key)
	} else {
		logger.Printf("[cache] MISS key=%s\n", key)
	}
	return value, ok
}

// Set é exclusivo: bloqueia leituras e escritas
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	logger.Printf("[cache] SET key=%s\n", key)
	c.data[key] = value
}
