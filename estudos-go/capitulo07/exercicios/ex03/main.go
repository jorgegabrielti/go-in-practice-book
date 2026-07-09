// O Marcador de Duplicadas (Desafio)
// Dado o slice: licao := []int{2,5,2,8,5,9,2}.
// Escreva um algoritmo que crie um novo slice contendo apenas os números únicos (sem repetição).
// Resultado esperado: [2 5 8 9]. Dica: você vai precisar de um loop dentro de outro loop (para verificar se já existe), ou pode esperar até aprendermos sobre MAPS no próximo capítulo para fazer isso de forma eficiente.
// Por enquanto, tente fazer com loops ("força bruta")

package main

import "fmt"

func main() {

	licao := []int{2, 5, 2, 8, 5, 9, 2}

	unicos := []int{}

	for _, valor := range licao { // Na 1º iteraçao, o valor é 2.
		jaExiste := false // A variável 'jaExiste' é declarada e inicializada com 'false'.
		//Ela serve como uma "bandeirinha" para sabermos se o número já foi visto antes.

		// Verifica se o valor já está no slice 'unicos'
		for _, existente := range unicos { // Aqui vai verificar se o valor já existe no slice 'unicos'. A 1º iteraçao será com o valor 2. Como 'unicos' está vazio, o loop não será executado.
			if existente == valor {
				jaExiste = true
				break // Sai do loop interno, não adianta procurar mais
			}
		}

		// Só adiciona se NÃO existir
		if !jaExiste { // Como 'jaExiste' é false, ele entra no if e adiciona o valor ao slice 'unicos'. O resultado agora é [2].
			unicos = append(unicos, valor) // 1º iteraçao: 'unicos' agora é [2] | 2º iteraçao 'unicos' será [2 5] |
		}
	}
	fmt.Println("Lista de unicos:", unicos)
}
