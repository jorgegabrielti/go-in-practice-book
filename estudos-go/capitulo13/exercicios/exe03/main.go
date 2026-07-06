package main

// Exercício 3: A Pilha de Pratos (defer) — Desafio Conceitual
//
// O defer funciona como uma pilha (LIFO - Last In, First Out).
// O último a entrar é o primeiro a sair.
//
// Escreva um loop for de 0 a 4. Dentro do loop, coloque defer fmt.Println(i).
// Imprima "Fim do Loop". Rode o programa e observe a ordem dos números.
// Explique por que a contagem sai invertida (4, 3, 2, 1, 0).

import "fmt"

func main() {
	fmt.Println("Exercício 3: A Pilha de Pratos")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Fim do Loop")
}
