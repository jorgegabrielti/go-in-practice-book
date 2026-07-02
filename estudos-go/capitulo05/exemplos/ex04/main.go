package main

import "fmt"

func main() {
	for i := 0; i < 1000; i++ {
		if i == 5 {
			fmt.Println("Achei o número 5! Parando a busca.")
			break // Sai do loop. Não vai até 999.
		}
	}
}
