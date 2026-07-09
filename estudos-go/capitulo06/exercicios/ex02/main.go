package main

import "fmt"

func analisarPreco(preco float64) (bool, string) {
	if preco >= 100.0 {
		return true, "Caro"
	}
	return false, "Barato"
}

func main() {
	fmt.Println(analisarPreco(50.0))
	fmt.Println(analisarPreco(100.0))
}
