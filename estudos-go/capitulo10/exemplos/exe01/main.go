package main

import "fmt"

// Definição do Molde
type Produto struct {
	Nome string
	Preco float64
	EmEstoque int
	Ativo bool
}

// Struct aninhada
type Pedido struct {
	ID int
	Cliente string
	Itens []Produto // Um slice de Produtos!
}

func main() {
	//1. Criando um produto
	p1 := Produto{
		Nome: "Teclado Mecânico",
		Preco: 250.00,
		EmEstoque: 10,
		Ativo: true,
	}

//2. Imprimindo (O %+v é um truque para ver os nomes dos campos)
fmt.Printf("Produto Detalhado: %+v\n", p1)

//3. Ponteiro para Struct(Muito comum para evitar cópia)
//Vamos criar um ponteiro p2 que aponta para o p1
p2 := &p1
// Em C, teríamos que usar (*p2).Preco.
// Em Go, o compilador é gentil e deixa usar p2.Preco direto.
p2.Preco = 200.00 // Isso altera o p1 original

fmt.Println("Preço alterado no original: ", p1.Preco)

//4. Pedido com Lista de Itens
pedido := Pedido {
	ID: 1001,
	Cliente: "Roberto Justus",
	Itens: []Produto{
		p1,
		{Nome: "Mouse", Preco: 50.0}, // Criando struct inline
	},
}

// Calculando total do pedido
total := 0.0
for _, item := range pedido.Itens{
	total += item.Preco
}

fmt.Printf("Total do pedido %d: R$ %.2f\n", pedido.ID, total)

}
