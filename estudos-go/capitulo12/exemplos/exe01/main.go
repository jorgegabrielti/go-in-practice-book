package main

import (
	"fmt"
	"math"
)

// O Contrato
type Geometrico interface {
	Area() float64
	Perimetro() float64
}

// Tipo 1: Retângulo
type Retangulo struct {
	Largura, Altura float64
}

//Implementação Geometrico para Retangulo
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (r Retangulo) Perimetro() float64 {
	return 2*r.Largura + 2*r.Altura
}

// Tipo 2: Círculo
type Circulo struct {
	Raio float64
}

// Implementando Geometrico para Circulo
func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Raio
}

// Função Genérica que aceita o Contrato
func ExibirDetalhes(g Geometrico) {
	fmt.Println("--- Forma Geométria ---")
	fmt.Printf("Área: %.2f\n", g.Area())
	fmt.Printf("Perímetro: %.2f\n", g.Perimetro())

	// Type Assertion para fazer algo específico
	// "Se for um Círculo, mostre o raio"
	if c, ok := g.(Circulo); ok {
		fmt.Printf("Raio: %.2f(Específico de Círculo)\n", c.Raio)
	}
}

func main() {
	r := Retangulo{Largura: 10, Altura: 5}
	c := Circulo{Raio: 3}

	// O compilador aceita ambos, pois ambos implementam Area e Perímetro
	ExibirDetalhes(r)
	ExibirDetalhes(c)

	// Exemplo de Interface Vazia
	var listaGenerica []any
	listaGenerica = append(listaGenerica, 10)
	listaGenerica = append(listaGenerica, "Olá")
	listaGenerica = append(listaGenerica, true)

	fmt.Println("\nLista Genérica: ", listaGenerica)
}