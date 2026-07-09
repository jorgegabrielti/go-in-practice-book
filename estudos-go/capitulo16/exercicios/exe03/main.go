package main

// Exercício 3: O Temporizador Manual (Desafio)
//
// Crie uma função que aceita um canal. Ela dorme 2 segundos e depois envia true no canal.
// Na main, imprima "Esperando...". Faça uma operação de leitura no canal (<-c).
// Isso vai bloquear a main. Quando desbloquear, imprima "Pronto!".
// Você acabou de recriar o time.Sleep usando canais.

import (
	"fmt"
	"time"
)

// O temporizador aceita um canal de escrita (chan<- bool)
func temporizador(c chan<- bool) {
	// Dorme 2 segundos
	time.Sleep(2 * time.Second)
	// Envia true para sinalizar que o tempo acabou
	c <- true
}

func main() {
	fmt.Println("Exercício 3: O Temporizador Manual")
	c := make(chan bool)

	fmt.Println("Disparando temporizador...")
	// Lança a goroutine do temporizador
	go temporizador(c)

	fmt.Println("Esperando...")
	start := time.Now()

	// Operação de leitura que bloqueará a main até que o temporizador envie o valor
	<-c

	duracao := time.Since(start)
	fmt.Printf("Pronto! (Tempo de espera: %s)\n", duracao)
}
