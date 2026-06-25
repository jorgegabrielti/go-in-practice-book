// O mistério do Zero (Desafio)
// Declare três variáveis usando var, mase sem dar valor a elas e imprima as três. Para o texto, como ele é vazio, talvez não pareça nada no terminal. Desafio Extra: Pesquise como usar o comando fmt.Printf("%q", texto) para fazer o Go mostrar as aspas vazias ""explicitamente. O %q é um "verbo de formatação" quoted muito útil para debug. Tente descobrir o que ele faz

package main

import "fmt"

var numero int
var texto string
var boleano bool

func main() {
	fmt.Printf("%d %q %t\n", numero, texto, boleano)
}
