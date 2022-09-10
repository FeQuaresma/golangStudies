package main

import "fmt"

func main() {

	retorno := func(texto string) string {

		return fmt.Sprintf("Recebido -> %s", texto)
	}("Ola Mundo!")

	fmt.Println(retorno)
}
