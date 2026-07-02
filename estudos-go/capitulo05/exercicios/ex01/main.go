// Escolha um número (Ex: 7). Use um loop for clássico para imprimir a tabuada desse número de 1 a 10. Saída esperada: 7 x 1 = 7, 7 x 2 = 14...

package main

import "fmt"

func main() {
	numero := 7
	for i := 1; i <= 10; i++ {
		resultado := numero * i
		fmt.Printf("%d x %d = %d\n", numero, i, resultado)
	}
}
