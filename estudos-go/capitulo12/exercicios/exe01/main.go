//1. A Impressora (Fácil)
// Crie uma interface Imprimivel com um método Imprimir() string.
// Crie duas structs: Livro (Titulo, Autor) e Carro (Modelo, Marca). Implemente o método Imprimir para ambas (retorne uma frase bonitinha com os dados).
// Na main, crie um slice []Imprimivel, coloque um livro e um carro dentro, e use um range para imprimir tudo.

package main

import "fmt"

type Imprimivel interface {
	Imprimir() string
}

type Livro struct {
	Titulo string
	Autor  string
}

func (l Livro) Imprimir() string {
	return fmt.Sprintf("Livro: %s por %s", l.Titulo, l.Autor)
}

type Carro struct {
	Modelo string
	Marca  string
}

func (c Carro) Imprimir() string {
	return fmt.Sprintf("Carro: %s %s", c.Marca, c.Modelo)
}

func main() {
	var lista []Imprimivel
	lista = append(lista, Livro{Titulo: "O Senhor dos Anéis", Autor: "J.R.R. Tolkien"})
	lista = append(lista, Carro{Marca: "Toyota", Modelo: "Corolla"})

	for _, item := range lista {
		fmt.Println(item.Imprimir())
	}
}
