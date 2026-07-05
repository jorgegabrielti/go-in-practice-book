//1. O cadastro de Gamer (Fácil)
// Crie uma struct 'Jogador' com:
// - Nickname (string)
// - Nivel(int)
// Vivo (bool) Instancie um jogador chamado "NoobMaster", nível 1, Vivo=true. Imprima os dados.

package main

import "fmt"

type Jogador struct {
	NickName string
	Nivel    int
	Vivo     bool
}

func main() {
	var jogador1 Jogador

	jogador1.NickName = "NoobMaster"
	jogador1.Nivel = 1
	jogador1.Vivo = true

	fmt.Println("Nickname:", jogador1.NickName)
	fmt.Println("Nivel:", jogador1.Nivel)
	fmt.Println("Vivo:", jogador1.Vivo)

	fmt.Printf("%+v\n", jogador1)
}