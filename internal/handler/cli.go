package handler

import (
	"calculator/internal/domain"
	"fmt"
)

func RunCalculator(
	add domain.Operation,
	subtract domain.Operation,
	multiply domain.Operation,
	divide domain.Operation,
) {
	var x, y float64

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&x)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&y)

	fmt.Println("\nResultados:")

	if res, err := add.Calculate(x, y); err == nil {
		fmt.Printf("Soma: %.2f\n", res)
	}

	if res, err := subtract.Calculate(x, y); err == nil {
		fmt.Printf("Subtração: %.2f\n", res)
	}

	if res, err := multiply.Calculate(x, y); err == nil {
		fmt.Printf("Multiplicação: %.2f\n", res)
	}

	if res, err := divide.Calculate(x, y); err == nil {
		fmt.Printf("Divisão: %.2f\n", res)
	} else {
		fmt.Println("Erro na divisão:", err)
	}
}
