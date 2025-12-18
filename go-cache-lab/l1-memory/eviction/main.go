package main

func main() {
	logger.Println("[app] starting application")

	cache := NewCache(2)

	cache.Set("a", "1")
	cache.Set("b", "2")

	cache.Get("a") // a vira o mais recente

	cache.Set("c", "3") // b é removido

	cache.Get("b") // MISS
	cache.Get("a") // HIT
	cache.Get("c") // HIT

	logger.Println("[app] finished")
}
