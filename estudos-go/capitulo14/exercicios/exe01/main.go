package main

// Exercício 1: A Organização da Geometria (Fácil)
//
// Crie uma pasta geometria/. Dentro dela, crie formas.go com package geometria.
// Crie uma função AreaQuadrado(lado float64) float64.
// Na main.go (fora da pasta), importe o pacote e calcule a área de um quadrado de lado 5.

import (
	"fmt"

	"go-in-practice-book/estudos-go/capitulo14/exercicios/exe01/geometria"
)

func main() {
	fmt.Println("Exercício 1: A Organização da Geometria")
	lado := 5.0
	area := geometria.AreaQuadrado(lado)
	fmt.Printf("A área do quadrado com lado %.2f é: %.2f\n", lado, area)
}
