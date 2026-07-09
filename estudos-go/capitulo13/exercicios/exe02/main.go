package main

// Exercício 2: O Validador de Senha (Médio)
//
// Crie uma função ValidarSenha(senha string) error. Regras:
// - Se tamanho < 8: Retorne erro "senha muito curta: tem apenas N caracteres" (use fmt.Errorf).
// - Se não tiver números (desafio): Retorne "senha precisa de número".
// - Se passar: Retorne nil.

import (
	"errors"
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("Exercício 2: O Validador de Senha")

	testes := []string{
		"12345",       // Curta
		"senhasegura", // Sem número
		"senha12345",  // Válida
	}

	for _, senha := range testes {
		err := validarSenha(senha)
		if err != nil {
			fmt.Printf("Senha: %-12s | Erro: %v\n", senha, err)
		} else {
			fmt.Printf("Senha: %-12s | OK (Válida)\n", senha)
		}
	}
}

func validarSenha(senha string) error {
	if len(senha) < 8 {
		return fmt.Errorf("senha muito curta: tem apenas %d caracteres", len(senha))
	}
	if !temNumero(senha) {
		return errors.New("senha precisa de número")
	}
	return nil
}

func temNumero(senha string) bool {
	for _, char := range senha {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}
