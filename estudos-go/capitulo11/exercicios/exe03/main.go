//3. A calculadora Orientada a Métodos (Conceito)
// Em vez de funções soltas soma(a,b), crie uma struct Calculadora
// que mantém um histórico. type Calculadora struct {Resultado float64} Métodos:
// - Somar(v float64): soma ao Resultado interno.
// - Subtrair(v float64): subtrai do Resultado interno.
// - Limpar() -> zera. Na main: calc.Somar(10); calc.Subtrair(2) imprima o calc.Resultado final (deve ser 13)

package main

import "fmt"

type Calculadora struct {
	Resultado float64
}

func (c *Calculadora) Somar(v float64) {
	c.Resultado += v
}

func (c *Calculadora) Subtrair(v float64) {
	c.Resultado -= v
}

func (c *Calculadora) Limpar() {
	c.Resultado = 0
}

func main() {
	calculadora := Calculadora{}
	calculadora.Somar(15)
	calculadora.Subtrair(2)
	fmt.Println("Resultado final:", calculadora.Resultado)
}
