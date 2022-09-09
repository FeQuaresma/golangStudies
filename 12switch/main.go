package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabádo"
	default:
		return "Número Inválido"
	}
}

func diaDaSemana2(numero int) string {
	var varDia string
	switch {
	case numero == 1:
		varDia = "Domingo"
		fallthrough
	case numero == 2:
		varDia = "Segunda-Feira"
	case numero == 3:
		varDia = "Terça-Feira"
	case numero == 4:
		varDia = "Quarta-Feira"
	case numero == 5:
		varDia = "Quinta-Feira"
	case numero == 6:
		varDia = "Sexta-Feira"
	case numero == 7:
		varDia = "Sabádo"
	default:
		varDia = "Número Inválido"
	}

	return varDia
}

func main() {
	fmt.Println("Switchs")

	dia := diaDaSemana(-7)
	fmt.Println(dia)

	fmt.Println("-=-=-=-=-=-")

	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)

}
