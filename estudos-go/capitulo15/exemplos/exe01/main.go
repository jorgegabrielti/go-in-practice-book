package main

import (
	"fmt"
	"runtime" // Vamos usar para ver quantos CPUs temos
	"time"
)

func tarefaPesada(id int) {
	fmt.Printf("Robô %d: Iniciando trabalho...\n", id)
	time.Sleep(2 * time.Second) // Simula trabalho duro
	fmt.Printf("Robô %d: Terminou!\n", id)
}

func main() {
	// Quantos corações (núcleos) esse computador tem?
	numCpu := runtime.NumCPU()
	fmt.Printf("Seu computador tem %d núcleos lógicos.\n", numCpu)

	// Por padrão, o Go tenta usar todos eles.
	// Você pode limitar com runtime.GOMAXPROCS(1) se quiser simular um PC antigo.

	fmt.Println("\n--- Começando a Linha de Montagem ---")
	start := time.Now()

	// Lançando 10 robôs.
	// Se fosse sequencial, levaria 20 segundos (10 * 2s).
	// Como é concorrente, deve levar apenas ~2 segundos (todos trabalham juntos).

	for i := 1; i <= 10; i++ {
		go tarefaPesada(i)
	}

	// O Hack do Sleep (Prometo que no Cap 18 ensinarei o jeito certo!)
	fmt.Println("O Chefe (Main) está esperando...")
	time.Sleep(3 * time.Second)

	duracao := time.Since(start)
	fmt.Printf("\nTempo total: %s (Impressionante, não?)\n", duracao)
}
