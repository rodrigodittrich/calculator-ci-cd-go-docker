package main

import (
	"calculator/internal/handler"
	"calculator/internal/usecase"
)

func main() {
	handler.RunCalculator(
		usecase.Add{},
		usecase.Subtract{},
		usecase.Multiply{},
		usecase.Divide{},
	)
}
