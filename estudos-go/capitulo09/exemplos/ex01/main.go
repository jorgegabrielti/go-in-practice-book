package main

import "fmt"

func main() {
	idade := 30
	endereco := &idade
	fmt.Println("Endereço de memória (Hexadecimal): ", endereco)
	fmt.Println("Valor: ", *endereco)
	// Mudando o valor na memória
	*endereco = 40
	fmt.Println("Valor: ", *endereco)
}
