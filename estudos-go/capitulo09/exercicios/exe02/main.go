//2. O incrementador
// Crie uma variável contador := 0. Crie uma função incrementar que aceite um ponteiro e adicione +1 ao valor. Faça um loop que chame
// essa função 10 vezes passando o endereço do contador.
// Imprima o resulta final (deve ser 10).

package main

import "fmt"

func incrementar(numero *int) {
	*numero = *numero + 1
}

func main() {
	contador := 0

	for i := 0; i < 10; i++ {
		incrementar(&contador)
	}

	fmt.Println(contador)
}
