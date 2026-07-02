package main

import (
	"fmt"
	"time"
)

func main() {

	//1. If clássico
	numero := 42
	if numero%2 == 0 { // O operador % pega o resto da divisão(Módulo)
		fmt.Println(numero, "é Par")
	} else {
		fmt.Println(numero, "é ímpar")
	}

	//2. If com Inicialização Curta(Short Statement)
	// 'x' só existe dentro deste bloco 'if'
	if x := 10 * 2; x > 15 {
		fmt.Println("X é grande: ", x)
	}

	//fmt.Println(x) // Se descomentar, dá erro de "undefined: x"

	//3. Switch Clássico
	sabor := "Chocolate"

	switch sabor {
	case "chocolate":
		fmt.Println("Meu favorito!")
	case "baunilha":
		fmt.Println("Clássico!")
	case "morango", "flocos":
		fmt.Println("Frutado!")
	default:
		fmt.Println("Hmm, interessante.")
	}

	//4. Switch sem expressão
	// Equivalente a: switch true {
	// O Go verifica cada caso em ordem
	hoje := time.Now().Weekday()
	fmt.Println("Hoje é: ", hoje)

	switch {
	case hoje == time.Saturday || hoje == time.Sunday:
		fmt.Println("É fim de semana!")
	case hoje == time.Monday:
		fmt.Println("Força, guerreiro. É segunda!")
	default:
		fmt.Println("Bora trabalhar!")
	}

	//5. O polêmico "fallthrough"
	// Em Go, o switch sai no primeiro case que encontra.
	// Se você quiser que ele continue descendo (comportamento estilo C),
	// precia usar fallthrough. É raro, mas existe.

	nivel := 1
	fmt.Println("\nNíveis de Acesso.")

	switch nivel {
	case 1:
		fmt.Println("- Acesso básico")
		fallthrough // Permite continuar para o próximo case
	case 2:
		fmt.Println("- Acesso moderado")
		fallthrough
	case 3:
		fmt.Println("- Acesso total")
	default:
		fmt.Println("Nível desconhecido")
	}

}
