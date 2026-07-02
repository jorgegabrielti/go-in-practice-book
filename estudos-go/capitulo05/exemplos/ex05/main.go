package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("Pulando o 2 (número do azar)...")
			continue // Volta para o topo, ignora o Println de baixo
		}
		fmt.Println("Processando número: ", i)
	}
}
