// Escreva uma função trocar(a, b *int) que recebe dois ponteiros para inteiros e troca os valores deles.
// Na main : x :=5, y := 10. Chame trocar(&x,&y).
//fmt.Println(x, y) deve mostrar 10 5.
// Dica: você vai precisar de uma variável temporária ou usar o truque de múltipla atribuição do Go: *a, *b = *b, *a

package main

import "fmt"

func trocar(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func main() {
	x := 5
	y := 10
	trocar(&x, &y)
	fmt.Println(x, y)
}