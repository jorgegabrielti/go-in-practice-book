package main

// Exercício 2: O Segredo (Médio)
//
// No pacote geometria, crie uma variável PRIVADA: numeroMagico = 42
// Crie uma função PÚBLICA: PegarSegredo() int que retorna numeroMagico
// Na main:
// - Tente acessar geometria.numeroMagico (deve causar erro de compilação)
// - Tente acessar geometria.PegarSegredo() (deve funcionar)

import (
	"fmt"

	"go-in-practice-book/estudos-go/capitulo14/exercicios/exe02/geometria"
)

func main() {
	fmt.Println("Exercício 2: O Segredo")

	// Tentar acessar diretamente o segredo causará erro de compilação:
	// fmt.Println(geometria.numeroMagico) // ERRO: cannot refer to unexported name geometria.numeroMagico

	// Acessando o segredo através da função pública:
	segredo := geometria.PegarSegredo()
	fmt.Printf("O segredo recuperado do pacote geometria é: %d\n", segredo)
}
