package main

import "fmt"

var sigla string = "He"

func main() {
	switch sigla {
	case "He":
		fmt.Println("He => Hélio")
	case "O":
		fmt.Println("O => Oxigênio")
	case "Li":
		fmt.Println("Li => Lítio")
	default:
		fmt.Println("Elemento desconhecido")
	}
}
