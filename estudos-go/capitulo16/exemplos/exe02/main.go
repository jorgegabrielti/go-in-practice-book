package main

import (
	"fmt"
	"time"
)

// O Gerador (Produtor)
// Só envia dados (chan<- int)
func gerarNumeros(c chan<- int) {
	fmt.Println("Gerador: Iniciando produção...")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Gerador: Enviando %d...\n", i)
		c <- i
		time.Sleep(200 * time.Millisecond) // Simula algum atraso na produção
	}
	fmt.Println("Gerador: Produção finalizada. Fechando canal.")
	close(c)
}

// O Consumidor
// Só recebe dados (<-chan int)
func calcularQuadrado(c <-chan int) {
	fmt.Println("Consumidor: Iniciando processamento...")
	for numero := range c {
		quadrado := numero * numero
		fmt.Printf("Consumidor: Recebi %d -> Quadrado = %d\n", numero, quadrado)
	}
	fmt.Println("Consumidor: Canal fechado. Processamento finalizado.")
}

func main() {
	fmt.Println("Iniciando Exemplo 2: Gerador e Consumidor")
	esteira := make(chan int)

	// Lança o gerador em background
	go gerarNumeros(esteira) // Vai travar a cada envio até o consumidor ler

	// O Consumidor roda na Main
	calcularQuadrado(esteira)

	// Como calcularQuadrado usa um loop que só sai quando o canal fecha,
	// a main espera tudo terminar naturalmente. Sem time.Sleep gambiarras!
	fmt.Println("Fim do Exemplo 2!")
}
