package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo da main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("felipequada@gmail.com")
	fmt.Println(erro)

}
