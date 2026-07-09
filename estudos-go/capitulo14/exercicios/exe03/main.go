package main

// Exercício 3: A Biblioteca Externa (Prático)
//
// 1. No terminal: go get github.com/google/uuid
// 2. Importe a biblioteca no código
// 3. Chame uuid.New().String() e imprima o resultado
// 4. Rode e veja um ID do tipo: f47ac10b-58cc-4372-a567-0e02b2c3d479

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Exercício 3: A Biblioteca Externa")

	// Gerando e imprimindo um UUID aleatório
	id := uuid.New().String()
	fmt.Printf("UUID gerado com sucesso: %s\n", id)
}
