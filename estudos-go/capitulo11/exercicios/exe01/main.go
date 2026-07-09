//1. O Carro Acelerado (Fácil)
// Crie uma struct Carro com campo VelocidadeAtual (int).
// Crie um método Acelerar() que aumente a velocidade em 10. Crie um método Frear() que diminua em 10(mas não deixe ficar negativo).
// Na main, intancie o carro, acelere 3 vezes e frei 1 vez. Imprima a velocidade final.

package main

import "fmt"

type Carro struct {
	VelocidadeAtual int
}

func (c *Carro) Acelerar() {
	c.VelocidadeAtual += 10
}

func (c *Carro) Frear() {
	if c.VelocidadeAtual < 10 {
		c.VelocidadeAtual = 0
	} else {
		c.VelocidadeAtual -= 10
	}
}

func main() {
	carro := Carro{}
	carro.Acelerar()
	carro.Acelerar()
	carro.Acelerar()
	carro.Frear()
	fmt.Println("Velocidade final:", carro.VelocidadeAtual)
}
