//1. O Contador de Palavras (Clássico médio)
// Imagine que você tem uma frase(string): "banana maçã laranja maçã banana".
// Você quer saber quantas vezes cada palavra aparece. O resultado deve ser um map: {
	// "banana": 3,
	// "maçã": 2,
	// "laranja": 1
// }
// Dica: você pode criar um slice de palavras mmanualmente []string{"banana", ...} e percorrer ele, somando +1 no map para cada palavra encontrada. 

package main

import "fmt"

func main() {
	palavras := []string{"banana", "maçã", "laranja", "maçã", "banana"}

	//Criando o map para armazenar a contagem
	contador := make(map[string]int)

	//2. Iterando sobre o slice de palavras e somando +1 no map
	for _, palavra := range palavras {
		contador[palavra]++ // O valor +1 sempre será adicionado ao valor inicial do map que é 0.
		//contador[palavra] = contador[palavra] +1
	}

	fmt.Println("Contagem de palavras:", contador)

	//2. Exibindo os resultados ordenados alfabeticamente por chave (nome da fruta)
	
}
