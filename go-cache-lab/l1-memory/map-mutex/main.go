package main

import (
	"fmt"
	"sync"
)

// main é o ponto de entrada da aplicação.
// Aqui simulamos múltiplas goroutines acessando o cache
// ao mesmo tempo.
func main() {
	logger.Println("[app] starting application")

	// Cria o cache em memória
	cache := NewCache()

	// WaitGroup garante que o programa só finalize
	// após todas as goroutines terminarem.
	var wg sync.WaitGroup

	// Dispara várias goroutines simulando acesso concorrente
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			// Reutilizamos poucas chaves de propósito
			// para gerar contenção e acessos concorrentes.
			key := fmt.Sprintf("key-%d", i%3)

			// Escrita no cache
			cache.Set(key, "value")

			// Leitura no cache
			cache.Get(key)
		}(i)
	}

	// Aguarda todas as goroutines finalizarem
	wg.Wait()

	// Exibe o estado final do cache
	cache.Dump()

	logger.Println("[app] finished")
}
