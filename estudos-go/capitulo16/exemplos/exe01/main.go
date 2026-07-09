package main

import (
	"fmt"
	"time"
)

// --- Exemplo 1: O Ping Pong (Sincronização Perfeita) ---

// O Jogador — bloqueia esperando a bola, devolve, repete.
func jogarPingPong(nome string, c chan string) {
	for {
		// 1. Espera a bola chegar (Bloqueia aqui)
		bola := <-c
		fmt.Printf("%s: Recebi %s\n", nome, bola)
		time.Sleep(500 * time.Millisecond) // Efeito dramático

		// 2. Devolve a bola (Bloqueia até o outro pegar)
		fmt.Printf("%s: Devolvendo...\n", nome)
		c <- bola
	}
}

func exemplo1PingPong() {
	fmt.Println("=== Exemplo 1: Ping Pong ===")
	mesa := make(chan string)

	go jogarPingPong("Ping", mesa)
	go jogarPingPong("Pong", mesa)

	// Coloca a bola na mesa para começar
	mesa <- "bola"

	// Deixa os jogadores trocarem por 3 segundos
	time.Sleep(3 * time.Second)
	fmt.Println("\nJogo encerrado!")
}

// --- Exemplo 2: Gerador e Consumidor (for range + close) ---

// gerarNumeros envia números de 1 a 5 no canal e fecha.
func gerarNumeros(esteira chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Gerando: %d\n", i)
		esteira <- i // Bloqueia até calcularQuadrado ler
	}
	close(esteira) // Sinaliza que não há mais dados
}

// calcularQuadrado lê do canal até ele fechar.
func calcularQuadrado(esteira chan int) {
	for numero := range esteira { // for range sai sozinho quando o canal fecha
		fmt.Printf("Quadrado de %d = %d\n", numero, numero*numero)
	}
}

func exemplo2GeradorConsumidor() {
	fmt.Println("\n=== Exemplo 2: Gerador e Consumidor ===")
	esteira := make(chan int)

	// Lança o gerador em background
	go gerarNumeros(esteira)

	// O consumidor roda na main — bloqueia até o canal fechar
	calcularQuadrado(esteira)

	// Como calcularQuadrado usa for range que só sai quando o canal fecha,
	// a main espera tudo terminar naturalmente. Sem time.Sleep gambiarras!
	fmt.Println("Processamento concluído.")
}

func main() {
	exemplo1PingPong()
	exemplo2GeradorConsumidor()
}
