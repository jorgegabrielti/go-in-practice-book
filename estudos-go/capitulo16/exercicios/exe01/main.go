package main

// Exercício 1: O Correio Elegante (Fácil)
//
// Crie um canal de strings. Crie uma goroutine que envia "Olá", "Mundo", "Do", "Go".
// Na main, receba e imprima essas 4 mensagens manualmente usando 4 fmt.Println(<-c).
// O que acontece se você tentar imprimir uma 5ª vez? (Deadlock).

import (
	"fmt"
	"time"
)

func enviarMensagens(c chan string) {
	mensagens := []string{"Olá", "Mundo", "Do", "Go"}
	for _, msg := range mensagens {
		fmt.Printf("[Goroutine] Enviando: %s\n", msg)
		c <- msg
		time.Sleep(100 * time.Millisecond) // Pequena pausa
	}
}

func main() {
	fmt.Println("Exercício 1: O Correio Elegante")
	c := make(chan string)

	// Lançando a goroutine
	go enviarMensagens(c)

	// Recebendo e imprimindo manualmente 4 vezes
	fmt.Println("Recebido 1:", <-c)
	fmt.Println("Recebido 2:", <-c)
	fmt.Println("Recebido 3:", <-c)
	fmt.Println("Recebido 4:", <-c)

	// Se tentarmos receber uma 5ª vez:
	// O canal estará vazio e a goroutine enviarMensagens já terminou de executar.
	// O Go runtime detectará que a goroutine main está bloqueada aguardando
	// dados em um canal de onde ninguém mais pode enviar, resultando em:
	// "fatal error: all goroutines are asleep - deadlock!"
	//
	// Experimente descomentar a linha abaixo para ver o deadlock:
	// fmt.Println("Recebido 5:", <-c)

	fmt.Println("\nSucesso! O canal foi lido 4 vezes e o programa terminou sem deadlocks.")
}
