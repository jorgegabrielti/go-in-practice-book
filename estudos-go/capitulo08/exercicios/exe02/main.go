//2. O dicionário de  Cores (Fácil)
// Crie um map onde a chave é o nome da cor em Português e o Valor é em inglês.
// Exemplo : {"Vermelho" : "Red"}
// Peça para uma variável corBuscada := "Azul".
// Use o comma ok para verificar se existe. Se existir, imprima a tradução. Se não, imprima "Cor não encontrada".

package main

import "fmt"

func main() {
	//1. Criando o map com cores
	cores := map[string]string{
		"Vermelho": "Red",
		"Verde":    "Green",
		"Amarelo":  "Yellow",
		"Laranja":  "Orange",
	}

	//2. Buscando uma cor
	corBuscada := "Vermelho"

	// COMMA OK IDIOM
	// Ao acessar um map com dois retornos, o Go devolve:
	//   - traducao → o valor associado à chave (ou "" se não existir)
	//   - ok       → true se a chave existe, false se não existe
	//
	// Sem o comma ok:
	//   traducao := cores[corBuscada]
	//   → se "Azul" não existir, traducao = "" — impossível saber se a chave
	//     realmente não existe ou se existia com valor vazio.
	//
	// Com o comma ok:
	//   cores["Vermelho"] → traducao = "Red", ok = true
	//   cores["Azul"]     → traducao = "",    ok = false  ← distinção clara
	traducao, ok := cores[corBuscada]

	if ok {
		fmt.Printf("A cor \"%s\" em Inglês é %s\n", corBuscada, traducao)
	} else {
		fmt.Printf("A cor \"%s\" não foi encontrada!\n", corBuscada)
	}
}
