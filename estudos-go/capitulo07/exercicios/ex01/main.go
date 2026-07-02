//1. A lista de Convidades (Fácil)
// Crie um Slice de strings vazio. Use um loop for (que roda 5 vezes) para pedir ao usuário digitar um nome(ou apenas simule nomes como "Convidade1", "Convidado 2"). A cada volta, faça append na lista. No final, imprima a lista completa e o tamanho dela.

package main

import "fmt"

func main(){
	convidados :=[]string{}
	for i:=0; i < 5; i++ {
		convidados = append(convidados, fmt.Sprintf("Convidado %d", i+1)) // O que faz exatamente o Sprintf?
	}
	fmt.Println("Lista de convidados: ", convidados)
	fmt.Println("Tamanho da lista: ", len(convidados))
}