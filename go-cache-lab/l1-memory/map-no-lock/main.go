package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("[app] starting application")

	// Usa todos os cores disponíveis
	runtime.GOMAXPROCS(runtime.NumCPU())

	// map SEM qualquer proteção
	cache := make(map[string]string)

	var wg sync.WaitGroup

	workers := 100
	iterations := 1000

	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			for j := 0; j < iterations; j++ {
				key := fmt.Sprintf("key-%d", j%5)

				// Escritas concorrentes no map
				cache[key] = fmt.Sprintf("worker-%d", id)

				// Leitura concorrente
				_ = cache[key]
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("[app] finished")
}
