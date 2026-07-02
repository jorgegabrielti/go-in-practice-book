// Comece pelo número grande, por exemplo n:= 12345. Use um loop for (estilo while) para somar todos os digitos desse núemro. Resultado esperado: 1 + 2 + 3 + 4 + 5 = 15.
// Pistas Matemáticas:
// n % 10 te dá o último digito (5)
// n / 10 (divisão inteira) remove o último digito (1234).
// Repita isso enquanto n > 0

package main

import "fmt"

func main() {
	numero := 12345
	soma := 0
	for numero > 0 {
		soma += numero % 10
		numero /= 10
	}

	fmt.Println("Soma dos digitos: ", soma)
}
