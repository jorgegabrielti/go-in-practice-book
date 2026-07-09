package main

import (
	"fmt"

	"go-in-practice-book/estudos-go/capitulo14/exemplos/exe01/conversor"
)

func main() {
	km := 100.0
	milhas := conversor.KmParaMilhas(km)

	fmt.Printf("%.2f Km é igual a %.2f Milhas\n", km, milhas)

	// Tente descomentar a linha abaixo — erro de compilação garantido:
	// conversor.milhasParaKm(50) // ERRO: cannot refer to unexported name
}
