package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

type usuarioFoda struct {
	usuario
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Quaresma"
	u.idade = 22
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	var usuario2 = usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2)

	var usuario3 = usuario{nome: "Daniel"}
	fmt.Println(usuario3.nome)
}
