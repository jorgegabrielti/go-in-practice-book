package main

import (
	"errors"
	"fmt"
)

// Erro customizado (sentinel error)
// É boa prática criar variáveis para erros conhecidos para poder comparar depois
var ErrDivisaoPorZero = errors.New("matemática diz não: divisão por zero")

// Função que pode falhar
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisaoPorZero
	}
	return a / b, nil
}

// Função simulando algo perigoso
func operacaoArriscada() {
	// Defer + Recover: O cinto de segurança
	defer func() {
		// recover() pega o valor do panic, se houver um
		if r := recover(); r != nil {
			fmt.Println("🚒 OPA! Recuperei de um pânico:", r)
		}
	}()

	fmt.Println("Iniciando operação nuclear...")
	panic("EXPLOSÃO NO REATOR 4!") // Simula um crash fatal
	fmt.Println("Isso nunca será impresso.")
}

func main() {
	// 1. Tratamento Padrão
	resultado, err := dividir(10, 0)

	if err != nil {
		// Podemos verificar se é um erro específico!
		if errors.Is(err, ErrDivisaoPorZero) {
			fmt.Println("Erro Específico: Você tentou dividir por zero.")
		} else {
			fmt.Println("Erro Genérico:", err)
		}
	} else {
		fmt.Println("Resultado:", resultado)
	}

	// 2. Testando Panic e Recover
	fmt.Println("\n--- Teste de Pânico ---")
	operacaoArriscada()
	fmt.Println("O programa continuou rodando! O recover salvou o dia.")

	// 3. Defer na Prática
	fmt.Println("\n--- Teste de Defer ---")
	// Função anônima auto-executável para isolar escopo
	func() {
		fmt.Println("Abrindo porta...")
		defer fmt.Println("Fechando porta (Defer).") // Vai pro final da fila

		fmt.Println("Entrando na casa...")
		fmt.Println("Fazendo bagunça...")
	}() // Fecha parênteses executa a função agora
}
