package main

import "fmt"

func main() {

	//--- 1. Arrays(fixos) ---
	var arrayFixo [3]int

	arrayFixo[0] = 10
	arrayFixo[1] = 20
	arrayFixo[2] = 30

	//arrayFixo[3] = 40 // ERRO DE COMPILAÇÃO: Índice fora dos limites

	fmt.Println("Array fixo: ", arrayFixo)

	// Array Literal(Já preenchido)
	vogais := [5]string{"A", "E", "I", "O", "U"}

	fmt.Printf("Vogais (Tamanho %d): %v\n", len(vogais), vogais)

	//--- 2. Slices(Dinâmicos) ---
	// Criando slice vazio

	var sliceDinamico []int
	fmt.Println("\nSlice Inicial: ", sliceDinamico)
	fmt.Println("Vazio?", sliceDinamico == nil)

	// Usando Append para crescer
	// Note que precisamos reatribuir: slice = append(slice, ...)
	sliceDinamico = append(sliceDinamico, 100)
	fmt.Println("Slice Crescido: ", sliceDinamico)
	sliceDinamico = append(sliceDinamico, 200, 300, 400) // Varios de uma vez
	fmt.Println("Slice Crescido: ", sliceDinamico)

	// --- 3. Capacidade vs Tamanho ---
	// Vamos observar o Go "trocando de array" quando enche
	numeros := make([]int, 0, 2) // Tamanho 0, Capacidade 2

	fmt.Printf("\nLen: %d, Cap: %d, Array: %v\n", len(numeros), cap(numeros), numeros)

	numeros = append(numeros, 1)
	numeros = append(numeros, 2)

	fmt.Printf("Len: %d, Cap: %d, Array: %v(Cheio!)\n", len(numeros), cap(numeros), numeros)

	numeros = append(numeros, 3) // Opa, Passou de 2. O Go precisa alocar mais memória

	// Geralmente ele dobra a capacidade para ter folga
	fmt.Printf("Len: %d, Cap: %d, Array: %v (Cresceu!)\n", len(numeros), cap(numeros), numeros)

	// --- 4. Slicing e Referência ---
	original := []string{"Batman", "Superman", "Mulher Maravilha"}

	copiaCopia := original[0:3]
	fmt.Println("\nOriginal: ", original)
	fmt.Println("Cópia (Slice): ", copiaCopia)

	// PERIGO: Mudando a cópia
	copiaCopia[0] = "Coringa"

	fmt.Println("--- Depois de mudar a cópia ---")
	fmt.Println("Cópia: ", copiaCopia)
	fmt.Println("Original: ", original) // O Batman virou Coringa no original também!

}
