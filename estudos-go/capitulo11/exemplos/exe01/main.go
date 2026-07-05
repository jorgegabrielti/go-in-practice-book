package main

import "fmt"

//Definindo a Struct
type ContaBancaria struct {
	Titular string
	Saldo   float64
}

// Método Value Receiver(Apenas lê o saldo, não muda nada)
// É como consultar o extrato.
func (c ContaBancaria) Extrato() {
	fmt.Printf("Conta de %s | Saldo: R$ %.2f\n", c.Titular, c.Saldo)
}

//Método Pointer Receiver (ALTERA o saldo)
// É como fazer um depósito.
func (c *ContaBancaria) Depositar(valor float64) {
	if valor > 0 {
		c.Saldo += valor
		fmt.Println("Depósito realizado com sucesso")
	} else {
		fmt.Println("Valor inválido para depósito")
	}
}

// Método Pointer Receiver (ALTERA o saldo)
// É como fazer um saque. Retorna erro se não tiver saldo.
func (c *ContaBancaria) Sacar(valor float64) bool {
	if c.Saldo >= valor {
		c.Saldo -= valor
		fmt.Println("Saque realizado.")
		return true
	}

	fmt.Println("Saldo insuficiente!")
	return false
}

func main() {
	//Criando a conta
	minhaConta := ContaBancaria{
		Titular: "César",
		Saldo:   100.00,
	}

	//1. Consultando (Value Receiver)
	minhaConta.Extrato()

	//2. Depositando (Pointer Receiver - Muda o saldo)
	minhaConta.Depositar(50.00)
	minhaConta.Extrato() // Deve ser 150.00

	//3. Sacando
	sucesso := minhaConta.Sacar(200.00) // Tenta sacar mais do que tem

	if !sucesso {
		fmt.Println("Preciso trabalhar mais...")
	}

	//Saldo continua 150
	minhaConta.Extrato()
}
