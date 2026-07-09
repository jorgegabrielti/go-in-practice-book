//3. O Type Switch (Desafio)
// Crie um função Classificar(v any). Use um switch v.(type) para imprimir
// - Se for string: "É um texto de tamanho X"
// - Se for int: "É umm número: dobro 2X"
// - Se for bool: "É booleano: valor invertido = !X"
// - Default: "Tipo desconhecido" Teste chamando a função com vários valores

package main

import "fmt"

func Classificar(v any) {
	switch val := v.(type) {
	case string:
		fmt.Printf("É um texto de tamanho %d\n", len(val))
	case int:
		fmt.Printf("É um número: dobro %d\n", val*2)
	case bool:
		fmt.Printf("É booleano: valor invertido = %t\n", !val)
	default:
		fmt.Println("Tipo desconhecido")
	}
}

func main() {
	Classificar("Olá, Go!")
	Classificar(10)
	Classificar(true)
	Classificar(3.14)
}
