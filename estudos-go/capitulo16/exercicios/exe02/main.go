package main

// Exercício 2: A Soma Distribuída (Médio)
//
// Crie uma função somar(valores []int, c chan int). Ela soma o slice e envia o total no canal.
// Na main:
// 1. Crie um slice grande: []int{1, 2, 3, 4, 5, 6}.
// 2. Divida ao meio: s1 := valores[:3] e s2 := valores[3:].
// 3. Lance duas goroutines: go somar(s1, c) e go somar(s2, c).
// 4. Receba os dois resultados (x, y := <-c, <-c) e some-os para o total final.

import (
	"fmt"
)

func somar(valores []int, c chan int) {
	soma := 0
	for _, v := range valores {
		soma += v
	}
	fmt.Printf("[Goroutine] Somando slice %v -> Resultado = %d\n", valores, soma)
	c <- soma
}

func main() {
	fmt.Println("Exercício 2: A Soma Distribuída")

	valores := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)

	// Dividindo o slice ao meio
	s1 := valores[:3] // [1, 2, 3]
	s2 := valores[3:] // [4, 5, 6]

	fmt.Printf("Slice original: %v\n", valores)
	fmt.Printf("Sub-slices: %v e %v\n\n", s1, s2)

	// Lançando as goroutines
	go somar(s1, c)
	go somar(s2, c)

	// Recebendo os dois resultados.
	// Nota: Como as goroutines rodam concorrentemente, o recebimento do canal
	// bloqueia a main até que ambos os valores sejam enviados.
	x := <-c
	y := <-c

	total := x + y
	fmt.Printf("\nResultado 1 recebido: %d\n", x)
	fmt.Printf("Resultado 2 recebido: %d\n", y)
	fmt.Printf("Soma total distribuída: %d (Esperado: 21)\n", total)
}
