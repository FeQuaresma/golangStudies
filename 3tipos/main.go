package main

import (
	"fmt"
)

func main() {
	var (
		num1   float64
		num2   float64
		opr    string
		result float64
	)

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o operador (+, -, *, /): ")
	fmt.Scan(&opr)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	switch opr {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2

	}

	fmt.Println("O resultado foi", result)

}
