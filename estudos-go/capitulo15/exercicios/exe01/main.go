package main

// Exercício 1: O Observador de CPUs (Fácil)
//
// Use runtime.NumCPU() e imprima quantos núcleos sua máquina tem.
// Depois, use runtime.GOMAXPROCS(1) para forçar o Go a usar apenas 1 núcleo.
// Rode o exemplo do main.go novamente. O tempo mudou?
// As mensagens dos robôs apareceram em ordem diferente?

import (
	"fmt"
	"runtime"
	"time"
)

func tarefaPesada(id int) {
	fmt.Printf("Robô %d: Iniciando trabalho...\n", id)
	time.Sleep(2 * time.Second) // Simula trabalho duro
	fmt.Printf("Robô %d: Terminou!\n", id)
}

func main() {
	fmt.Println("Exercício 1: O Observador de CPUs")

	// 1. Mostrar quantos núcleos lógicos o computador tem originalmente.
	numCpuOriginal := runtime.NumCPU()
	fmt.Printf("Seu computador tem originalmente %d núcleos lógicos.\n", numCpuOriginal)

	// 2. Limitar o runtime para usar apenas 1 núcleo.
	gOld := runtime.GOMAXPROCS(1)
	fmt.Printf("GOMAXPROCS alterado de %d para 1\n", gOld)

	fmt.Println("\n--- Começando a Linha de Montagem (com 1 CPU) ---")
	start := time.Now()

	// Lançando 10 robôs concorrentes.
	for i := 1; i <= 10; i++ {
		go tarefaPesada(i)
	}

	fmt.Println("O Chefe (Main) está esperando...")
	time.Sleep(3 * time.Second)

	duracao := time.Since(start)
	fmt.Printf("\nTempo total: %s\n", duracao)
	fmt.Println("\nRespostas às perguntas do Exercício:")
	fmt.Println("1. O tempo mudou?")
	fmt.Println("   Resposta: Não. O tempo de execução das tarefas continua em torno de ~2s (tempo de sleep das goroutines),")
	fmt.Println("   pois quando uma goroutine bloqueia (via time.Sleep), o Scheduler a remove da CPU/Thread única e")
	fmt.Println("   agenda outra que esteja pronta para rodar (multitasking cooperativo/preemptivo do Go Runtime).")
	fmt.Println("2. As mensagens dos robôs apareceram em ordem diferente?")
	fmt.Println("   Resposta: Sim. A ordem de inicialização e término ainda é não-determinística, pois depende do")
	fmt.Println("   agendador interno do Go gerenciar a fila de execução da thread única.")
}
