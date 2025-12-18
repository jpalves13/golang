package main

import (
	"container/list"
	"sync"
)

// entry representa um item do cache
type entry struct {
	key   string
	value string
}

type Cache struct {
	mu       sync.Mutex
	capacity int
	items    map[string]*list.Element
	order    *list.List
}

func NewCache(capacity int) *Cache {
	logger.Printf("[cache] initializing LRU cache (capacity=%d)\n", capacity)

	return &Cache{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		order:    list.New(),
	}
}

// Get move o item para o topo (mais recentemente usado)
func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		c.order.MoveToFront(elem)
		logger.Printf("[cache] HIT key=%s\n", key)
		return elem.Value.(*entry).value, true
	}

	logger.Printf("[cache] MISS key=%s\n", key)
	return "", false
}

// Set adiciona ou atualiza um item
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		// Atualiza valor e move para frente
		elem.Value.(*entry).value = value
		c.order.MoveToFront(elem)
		logger.Printf("[cache] UPDATE key=%s\n", key)
		return
	}

	// Eviction: remove o menos usado
	if c.order.Len() >= c.capacity {
		back := c.order.Back()
		if back != nil {
			ent := back.Value.(*entry)
			delete(c.items, ent.key)
			c.order.Remove(back)
			logger.Printf("[cache] EVICT key=%s\n", ent.key)
		}
	}

	elem := c.order.PushFront(&entry{key: key, value: value})
	c.items[key] = elem
	logger.Printf("[cache] SET key=%s\n", key)
}
