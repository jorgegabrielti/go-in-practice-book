package main

import "fmt"

func main() {
	nomes := []string{"Ana", "Beto", "Carla"}

	//for i, nome := range nomes {
	for _, nome := range nomes {
		fmt.Println(nome)
	}
}
