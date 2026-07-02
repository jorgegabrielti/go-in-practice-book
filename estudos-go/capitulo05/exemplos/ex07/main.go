package main

import (
	"fmt"
	"time"
)

func main() {

	//1. O For Clássico (Contagem Regressiva)
	fmt.Println("--- Contagem Regressiva ---")

	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("DECOLAR! ->")

	//2. O For como While(Acumulando some)
	fmt.Println("\n --- Somando até passar de 100 ---")

	soma := 0
	numero := 1
	for soma < 100 {
		soma = soma + numero
		fmt.Printf("Somando %d... Total: %d\n", numero, soma)

		numero++
	}

	//3. For infinito com Break(Jogo da Adivinhação Fake)
	fmt.Println("\n--- Tentando acertar a senha ---")
	tentativa := 0

	for {
		tentativa++
		// Simulando uma change de acertar
		if tentativa == 4 {
			fmt.Println("Senha correta encontrada na tentativa", tentativa)
			break
		}
		fmt.Println("Tentativa", tentativa, "falhou...")
	}

	//4. Usando 'continue'(Imprimir apenas ímpares)
	fmt.Println("\n --- Apenas ímpares ---")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println("\nFim!")
}
