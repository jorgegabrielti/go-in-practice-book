package main

import "fmt"

func main() {
	temperatura := 25

	if temperatura > 30 {
		fmt.Println("Ligando resfriamento máximo!")
	} else if temperatura > 20 {
		fmt.Println("Temperatura agradável")
	} else {
		fmt.Println("Ligando aquecedor")
	}
}
