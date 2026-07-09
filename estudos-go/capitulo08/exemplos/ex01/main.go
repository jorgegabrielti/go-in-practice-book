package main

import (
	"fmt"
	"sort"
)

func main() {
	//1. Criando e populando
	// Map de Aluno(String) -> Notas (Slice de Float)

	diario := make(map[string][]float64)

	diario["Carlos"] = []float64{8.5, 9.0}
	diario["Ana"] = []float64{10.0, 9.5, 8.0}
	diario["Beto"] = []float64{5.0}

	fmt.Println("Diário: ", diario)

	//2. Verificando Existência (Comma OK)
	aluno := "Pedro"

	notas, existe := diario[aluno] // Retorna o valor e um booleano indicando se a chave existe. Neste caso retorna um slide vazio e false.

	if existe {
		fmt.Println("Notas do Pedro: ", notas)
	} else {
		fmt.Println("O aluno ", aluno, "não está matriculado.")
	}

	//3. Iterando de forma aleatória
	fmt.Println("\n --- Lista de Chamada (Aleatória) ---")
	for nome, notas := range diario {
		fmt.Printf("Aluno: %s | Média: Calculando...\n", nome)
		//Pequeo cálculo de média aqui dentro.

		soma := 0.0

		for _, n := range notas {
			soma += n
		}

		fmt.Printf(" -> Média Final: %.2f\n", soma/float64(len(notas)))
	}
	//4. Imprimindo em Ordem Alfabética (O Jeito Certo)
	fmt.Println("\n --- Lista de Chamada (Ordenada) ---")
	//Passo A: Extrair as chaves

	var nomes []string
	for k := range diario {
		nomes = append(nomes, k)
	}

	//Passo B: Ordenar as chaves
	sort.Strings(nomes)

	//Passo C: Iterar na ordem
	for _, nome := range nomes {
		//Acessar o map direto
		fmt.Printf("%s -> %v\n", nome, diario[nome])
	}
}
