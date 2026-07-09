//2. A Comparação de Retângulos (Médio)
// Crie uma struct Retangulo com Largura e Altura. Crie uma função calcularArea(r Retangulo)int que retorna a área (L * A). No main, crie dois retângulos diferentes. Use um if para imprimir qual dos dois tem a maior área.
package main

import "fmt"

type Retangulo struct {
	Largura int
	Altura  int
}

func calcularArea(r Retangulo) int {
	return r.Largura * r.Altura
}

func main() {
	r1 := Retangulo{
		Largura: 10,
		Altura:  20,
	}
	r2 := Retangulo{
		Largura: 5,
		Altura:  15,
	}

	area1 := calcularArea(r1)
	area2 := calcularArea(r2)

	if area1 > area2 {
		fmt.Println("O retângulo 1 tem a maior área")
	} else {
		fmt.Println("O retângulo 2 tem a maior área")
	}
}
