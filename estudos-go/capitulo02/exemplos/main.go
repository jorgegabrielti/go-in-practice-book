package main

import "fmt"

var VersaoDoSistema = "1.0.0"

func main() {
	var totalDeProdutos int = 50
	fmt.Println("Total de Produtos(var):", totalDeProdutos)

	var estoqueReservado int
	fmt.Println("Estoque Reservado (Zero Value:)", estoqueReservado)

	produto := "Notebook Gamer"
	preco := 4500.99

	fmt.Println("Produto: ", produto, "Preço:", preco)

	preco = 4.200
	fmt.Println("Novo Preço:", preco)

	const TaxaDeEntrega = 15.00

	totalFinal := preco + TaxaDeEntrega
	fmt.Println("Total Final: ", totalFinal)

	var (
		nomeComprador  = "João Silva"
		emailComprador = "joao@email.com"

		ativo = true
	)

	fmt.Println("Cliente: ", nomeComprador, emailComprador, "Ativo?", ativo)
}
