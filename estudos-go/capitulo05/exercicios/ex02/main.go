// Este exercício é usado em entrevistas reais para filtar quem não sabe programar. Faça um loop de 1 a 50.
// Se o número for divisível por 3, imprima "Fizz".
// Se for divisível por 5, imprima "Buzz".
// Se for divisível por 3 e por 5, imprima "FizzBuzz".
// Senão, imprima apenas o número.

// Dica: A ordem das condições no if import! Verifique o "FizzBuzz" (15) antes dos outros.

package main

import "fmt"

func main() {
	for i := 1; i <= 50; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
