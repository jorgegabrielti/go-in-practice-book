package main

import "fmt"

func soma(a, b int) int { return a + b }

func subtracao(a, b int) int { return a - b }

func multiplicacao(a, b int) int { return a * b }

func main() {
	fmt.Println(soma(10, 5), subtracao(10, 5), multiplicacao(10, 5))
}
