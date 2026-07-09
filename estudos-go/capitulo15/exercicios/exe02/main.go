package main

// Exercício 2: O Loop Traiçoeiro (Médio)
//
// Escreva um laço for de 0 a 100. Dentro dele, lance uma goroutine anônima
// que imprime o número: go func() { fmt.Println(i) }()
// Rode o código. Você vê números repetidos? Você vê o número 100 ou 101?
// Descreva o caos. Depois conserte passando i como parâmetro da função anônima.

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Exercício 2: O Loop Traiçoeiro")

	// FASE 1: Demonstração do Caos / Gotcha de captura por referência
	// Para forçar o gotcha no Go 1.22+, declaramos a variável de controle do loop
	// fora da declaração do 'for', fazendo com que ela seja compartilhada entre
	// todas as iterações e closures das goroutines.
	fmt.Println("\n--- 1. Versão Caótica (Forçando o Gotcha) ---")
	var i int
	for i = 0; i <= 100; i++ {
		go func() {
			fmt.Printf("%d ", i)
		}()
	}

	// Aguardando um pouco para que todas as goroutines concorrentes executem.
	time.Sleep(1 * time.Second)
	fmt.Println()

	// FASE 2: Versão corrigida passando o parâmetro por valor
	fmt.Println("\n--- 2. Versão Corrigida (Passagem por Parâmetro) ---")
	for j := 0; j <= 100; j++ {
		go func(valor int) {
			fmt.Printf("%d ", valor)
		}(j)
	}

	// Aguardando a execução das goroutines da versão corrigida
	time.Sleep(1 * time.Second)
	fmt.Println()

	fmt.Println("\nRespostas às perguntas do Exercício:")
	fmt.Println("1. Você vê números repetidos?")
	fmt.Println("   Resposta: Sim, na versão caótica vemos muitos números repetidos (normalmente 101 aparecendo repetidamente).")
	fmt.Println("2. Você vê o número 100 ou 101?")
	fmt.Println("   Resposta: Vemos predominantemente o número 101 na versão caótica, porque o loop 'for' termina e define 'i = 101'")
	fmt.Println("   antes que a maioria das goroutines tenha tempo de ler e imprimir o valor de 'i'.")
	fmt.Println("3. Conserto:")
	fmt.Println("   Resposta: Passando o valor atual de cada iteração como argumento para a função anônima (go func(valor int) { ... }(j)),")
	fmt.Println("   fazendo com que cada goroutine receba uma cópia por valor do número naquele instante.")
}
