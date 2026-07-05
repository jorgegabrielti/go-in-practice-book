//2. O Processador de Pagamentos (Médio)
// Interface Pagamento com método Pagar(valor float64).
//Struct Boleto -> Método imprime "Pagando R$ X via Boleto". Struct Cartao -> Método imprime "Pagando R$ X via Cartão". Crie uma função FinalizarCompra(p Pagamento, valor float64).
// Teste pagando de formas diferentes.


package main

import "fmt"

// A Interface
type Pagamento interface {
	Pagar(valor float64) string
}

// Tipo 1: Boleto
type Boleto struct{}

func (b Boleto) Pagar(valor float64) string {
	return fmt.Sprintf("Pagando R$ %.2f via Boleto (Código de Barras será gerado)...")
}

// Tipo 2: Cartão
type Cartao struct {
	Numero string
}

func (c Cartao) Pagar(valor float64) string {
	return fmt.Sprintf("Pagando R$ %.2f via Cartão de Crédito (Final: %s)", valor, c.Numero)
}

// Função genérica
func FinalizarCompra(p Pagamento, valor float64) {
	fmt.Println(p.Pagar(valor))
}

func main() {
	b := Boleto{}
	c := Cartao{Numero: "1234"}

	valor := 100.0

	FinalizarCompra(b, valor)
	FinalizarCompra(c, valor)
}
