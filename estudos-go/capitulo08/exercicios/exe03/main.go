//3. O sistema de Votação (Desfio)
// Simule um urna eletrônica. Crie um map votos map[string]int.
// Crie uma função votar(candidato string) que incremente o voto no map. Simule vários votos: votar("Gopher"), votar("Python"), votar("Java"), votar("Python"), etc...
//No final, imprima quem ganhou (quem tem o maior int no map). Dica: Você vai precisar iterar sobre o mmapa e guardar numa variável vencedor quem tem o maior valor encontrado até agora.

package main

import "fmt"

var votos map[string]int

func votar(candidato string) {
	votos[candidato]++
}

func main() {
	votos = make(map[string]int)

	votar("Gopher")
	votar("Python")
	votar("Java")
	votar("Python")
	votar("Go")

	votosTotais := 0
	for _, total := range votos {
		votosTotais += total
	}
	fmt.Println("Total de votos: ", votosTotais)

	vencedor := ""
	maxVotos := -1

	for candidato, total := range votos {
		if total > maxVotos {
			maxVotos = total
			vencedor = candidato
		}
	}
	fmt.Println("Vencedor: ", vencedor)

}
