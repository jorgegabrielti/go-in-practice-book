package main

import "fmt"

// Função que recebe uma cópia(valor)

func zerarValor(x int) {
	x = 0 // Altera apenas a cópia local
}

// Função que recebe um endereço (ponteiro)

func zerarPonteiro(x *int) {
	*x = 0 // Vai até a memória original e zera
}

func main() {
	//1. Entendendo Endereços
	numero := 42
	var ponteiro *int = &numero

	fmt.Println("Valor de número: ", numero)
	fmt.Println("Endereço de número: ", ponteiro) // Algo como 0x0000...

	fmt.Println("O que tem no endereço: ", *ponteiro)

	//2. Mudando valor via ponteiro
	*ponteiro = 100 // Alterando indiretamente
	fmt.Println("Novo valor de número: ", numero) // Mágica! Mudou para 100

	//3. Teste Valor vs Referência
	a := 10
	b := 10

	zerarValor(a)
	zerarPonteiro(&b) // Passamos o endereço de b (&b)

	fmt.Println("A (Valor): ", a) // Continua 10
	fmt.Println("B(Ponteiro): ", b) // Virou 0

	//4. Ponteiro de Ponteiro (Inception) - Só por curiosidade
	// Ponteiro 'p' aponta para 'numero'
	// Ponteiro 'pp' aponta para ponteiro 'p'

	var pp **int = &ponteiro
	fmt.Println("Endereço do ponteiro: ", pp)
	fmt.Println("Valor lá no fundo: ", **pp) // 100
}