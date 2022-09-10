package main

import "fmt"

func fibonacci(posicao uint64) uint64 {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	// 1 1 2 3 4 5 8 13

	posicao := uint64(54)

	fmt.Println(fibonacci(posicao))
}
