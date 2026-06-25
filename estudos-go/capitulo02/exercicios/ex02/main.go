// Crie uma variável constante chamada PontoDeEbulição com o valor 100 (em Celsius). Crie um variável temperaturaAtual com o valor inicial 100. Coverta para Fahrenheit usando a fórmula: F = (C * 9/5) + 32. Imprima "A água ferve a 100ºC ou 212ºF". Dica: Você pode fazer a conta direto dentro do fmt.Println.

package main

import "fmt"

const PontoDeEbulicao = 100

var temperaturaAtual = 100

func main() {
	fmt.Println("A água ferve a ", PontoDeEbulicao, "Cº ou", (PontoDeEbulicao*1.8 + 32), "Fº")
}
