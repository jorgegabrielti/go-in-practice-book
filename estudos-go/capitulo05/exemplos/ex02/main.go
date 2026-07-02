package main

import "fmt"

func main() {
	energia := 100

	for energia > 0 {
		fmt.Println("Corendo... Energia: ", energia)
		energia = energia - 20
	}
	fmt.Println("Ufa! Cansei.")
}
