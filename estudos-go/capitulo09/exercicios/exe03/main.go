//3. O detetive de endereços (conceitual)
// Crie uma variável texto := "Go". Imprima o endereço dela (&texto). Chame uma função passando texto(por valor). Dentro da função, imprima o endereço do parâmetro recebido. Eles são iguais ou diferentes? Por quê?
// Lembre da analogia da foto.
// func verEndereco(t string) {
//	fmt.Println("Endereço dentro da função: ", &t)
//}

package main

import "fmt"

func verEndereco(t string) {
	fmt.Println("Endereço dentro da função: ", &t)
}

func main() {
	texto := "Go"
	fmt.Println("Endereço na main: ", &texto)
	verEndereco(texto)
}