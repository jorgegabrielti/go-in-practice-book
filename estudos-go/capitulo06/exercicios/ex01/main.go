package main

import "fmt"

func soma(a, b int)(int) {
  return a + b	
}

func subtracao(a, b int)(int) {
	return a - b
}

func multiplicacao(a, b int)(int) {
	return a * b
}

func main() {

	// Chamanda a função soma
	resultadoSoma := soma(10, 5)
	fmt.Printf("Soma %d + %d = %d\n", 10, 5, resultadoSoma)

	// Chamando a função subtração
	resultadoSubtracao := subtracao(10, 5)
	fmt.Printf("Subtracao %d - %d = %d\n", 10, 5, resultadoSubtracao)

	// Chamando a função multiplicação
	resultadoMultiplicacao := multiplicacao(10, 5)
	fmt.Printf("Multiplicacao %d * %d = %d\n", 10, 5, resultadoMultiplicacao)
}
