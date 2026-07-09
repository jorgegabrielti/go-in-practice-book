package main

// Exercício 3: O Ping-Pong Solitário (Desafio Conceitual)
//
// Crie uma função ping() que imprime "Ping" e dorme 500ms.
// Crie uma função pong() que imprime "Pong" e dorme 500ms.
// Lance as duas como goroutines. Faça a main dormir 5 segundos.
// Tente ver se eles ficam sincronizados ("Ping, Pong, Ping, Pong") ou se bagunça tudo.
// Spoiler: Sem canais (Cap 16), é impossível garantir a sincronia perfeita.

import (
	"fmt"
	"time"
)

func ping() {
	for {
		fmt.Println("Ping")
		time.Sleep(500 * time.Millisecond)
	}
}

func pong() {
	for {
		fmt.Println("Pong")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Exercício 3: O Ping-Pong Solitário")
	fmt.Println("Iniciando goroutines ping e pong...")

	go ping()
	go pong()

	// Mantém o main ativo por 5 segundos
	time.Sleep(5 * time.Second)

	fmt.Println("\nFim da execução da main.")
	fmt.Println("\nObservações conceituais:")
	fmt.Println("1. A ordem das mensagens Ping e Pong é garantida?")
	fmt.Println("   Resposta: Não. Embora às vezes pareça alternar perfeitamente, sob carga do sistema ou")
	fmt.Println("   ao longo do tempo, o agendador pode intercalar de forma imprevisível (ex: Ping Ping Pong ou Pong Pong Ping).")
	fmt.Println("   Como as duas goroutines agem de forma isolada, não há controle de dependência mútua.")
	fmt.Println("2. Como alcançar sincronia perfeita?")
	fmt.Println("   Resposta: Precisamos de canais (channels) ou outros mecanismos de sincronização (Capítulo 16)")
	fmt.Println("   para coordenar a troca de turnos entre as duas goroutines.")
}
