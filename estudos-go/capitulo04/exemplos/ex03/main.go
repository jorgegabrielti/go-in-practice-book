package main

import "fmt"

func main() {
	carteira := true
	sobrio := false

	if carteira && sobrio {
		fmt.Println("Pode dirigir")
	} else if carteira && !sobrio {
		fmt.Println("Você tem carteira, mas bebeu. Chame um Uber!")
	} else {
		fmt.Println("Você nem deveria estar perto do volante.")
	}
}
