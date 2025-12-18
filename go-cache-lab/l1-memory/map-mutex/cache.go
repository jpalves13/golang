package main

import "sync"

// Cache representa um cache em memória (L1)
// Ele usa um map para armazenar os dados e um Mutex
// para garantir acesso seguro em ambientes concorrentes.
type Cache struct {
	// mu protege o acesso ao map.
	// Sem ele, leituras e escritas concorrentes causariam panic.
	mu sync.Mutex

	// data armazena os valores do cache.
	// map NÃO é thread-safe em Go.
	data map[string]string
}

// NewCache cria uma nova instância de Cache.
// O map é inicializado vazio.
func NewCache() *Cache {
	logger.Println("[cache] initializing cache")

	return &Cache{
		data: make(map[string]string),
	}
}

// Get busca um valor no cache.
// O mutex é travado para garantir que nenhuma escrita
// ocorra ao mesmo tempo.
func (c *Cache) Get(key string) (string, bool) {
	// Lock bloqueia outras goroutines até que o acesso termine.
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.data[key]
	if ok {
		logger.Printf("[cache] HIT key=%s\n", key)
	} else {
		logger.Printf("[cache] MISS key=%s\n", key)
	}

	return value, ok
}

// Set insere ou atualiza um valor no cache.
// Também é protegido por mutex para evitar corrupção do map.
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	logger.Printf("[cache] SET key=%s\n", key)
	c.data[key] = value
}

// Dump imprime o estado atual do cache.
// Esse método é útil para fins didáticos e debug.
func (c *Cache) Dump() {
	c.mu.Lock()
	defer c.mu.Unlock()

	logger.Println("[cache] dumping cache state")
	for k, v := range c.data {
		logger.Printf("  %s => %s\n", k, v)
	}
}
