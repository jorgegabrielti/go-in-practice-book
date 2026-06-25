package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var pequeno uint8 = 255
	var grande int64 = 9223372036854775807

	fmt.Println("--- Inteiros ---")
	fmt.Println("Pequeno: %d (Tipo: %T)\n", pequeno, pequeno)

	fmt.Println("Grande: ", grande)

	var numero int = 10
	fmt.Println("Tamanho de 'int' neste computador: %d bytes\n", unsafe.Sizeof(numero))

	fmt.Println("\n--- Floats ---")

	var a float64 = 0.1
	var b float64 = 0.2
	var resultado = a + b

	fmt.Println("0.1 + 0.2 =", resultado)

	fmt.Println("É exatamente 0.3?", resultado == 0.3)

	//3.Strings e Runes

	fmt.Println("\n --- Strings ---")
	str := "Go é []"

	fmt.Println("String: ", str)

	fmt.Println("Tamanho em Bytes(len): ", len(str))

	//Para contar CARACTERES reais, convertemos para 'rune' slice

	runes := []rune(str)
	fmt.Println("Quantidade de caracteres (Runes):", len(runes))

	//4.Conversões(Casting)
	fmt.Println("\n --- Conversões ---")

	var nota1 int = 85
	var nota2 int = 90

	//Queremos a média. Se fizermos (nota1 + nota2)/2, será uma divisão inteira!

	//175/2 = 87 (perde 0.5)
	mediaErrada := (nota1 + nota2)/2
	fmt.Println("Média inteira (Errada):", mediaErrada)

	//Precisamos converter para float ANTES de dividir
	mediaCert := float64(nota1 + nota2) / 2.0
	fmt.Println("Média correta:", mediaCert)

}