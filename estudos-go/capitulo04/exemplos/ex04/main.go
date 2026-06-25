package main

import "fmt"

func main() {
	dia := "Sabado"

	switch dia {
	case "Segunda":
		fmt.Println("Dia de começar a semana!")
	case "Sexta":
		fmt.Println("Sextou!")
	case "Sabado", "Domingo":
		fmt.Println("Fim de semana!")
	default:
		fmt.Println("Bora trabalhar!")
	}
}
