package main

// Exercício 1: A Compra Segura (Fácil)
//
// Crie uma função Comprar(saldo, preco float64) (novoSaldo float64, err error).
// - Se preco > saldo: Retorne saldo original e um erro errors.New("saldo insuficiente").
// - Caso contrário: Retorne saldo - preco e nil.
// Teste na main com casos que dão certo e errado.

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Exercício 1: A Compra Segura")

	comprar(100, 50)
	comprar(50, 100)
}

func comprar(saldo, preco float64) (novoSaldo float64, err error) {
	if preco > saldo {
		return saldo, errors.New("saldo insuficiente")
	}
	return saldo - preco, nil
}
