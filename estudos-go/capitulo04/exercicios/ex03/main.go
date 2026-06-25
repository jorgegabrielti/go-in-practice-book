// Calculando ano bisexto

package main

import "fmt"

func main() {
	ano := 2026

	switch {
	case ano%4 == 0:
		fmt.Println("Ano bissexto")
	default:
		fmt.Println("Ano não bissexto")
	}
}
