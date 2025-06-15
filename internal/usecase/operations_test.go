package usecase

import (
	"testing"
)

func TestAdd(t *testing.T) {
	add := Add{}
	result, err := add.Calculate(10, 5)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	expected := 15.0
	if result != expected {
		t.Errorf("Esperado %.2f, obtido %.2f", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	sub := Subtract{}
	result, err := sub.Calculate(10, 5)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	expected := 5.0
	if result != expected {
		t.Errorf("Esperado %.2f, obtido %.2f", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	mult := Multiply{}
	result, err := mult.Calculate(10, 5)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	expected := 50.0
	if result != expected {
		t.Errorf("Esperado %.2f, obtido %.2f", expected, result)
	}
}

func TestDivide(t *testing.T) {
	div := Divide{}
	result, err := div.Calculate(10, 5)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	expected := 2.0
	if result != expected {
		t.Errorf("Esperado %.2f, obtido %.2f", expected, result)
	}
}

func TestDivideByZero(t *testing.T) {
	div := Divide{}
	_, err := div.Calculate(10, 0)
	if err == nil {
		t.Errorf("Esperava erro ao dividir por zero, mas não houve.")
	}
}
