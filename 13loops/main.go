package main

import (
	"fmt"
)

func main() {
	/* 	i := 0
	   	fmt.Println("Estrutura de Repetição")

	   	for i < 10 {
	   		time.Sleep(time.Second)
	   		i++
	   		fmt.Println(i)
	   	} */

	/* 	for j := 0; j < 10; j++ {
		time.Sleep(time.Second)
		fmt.Println("Incrementando J", j)
	} */

	nomes := [3]string{"Absoluto", "Daniel", "Quaresma"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	/* 	type usuarioStruct struct {
	   		nome      string
	   		sobrenome string
	   	}

	   	usuario2 := usuarioStruct{"Zé", "Roberto"}

	   	for chave, valor := range usuario2 {
	   		fmt.Println(chave, valor)
	   	} */

	for {
		fmt.Println("Infinito")
	}
}
