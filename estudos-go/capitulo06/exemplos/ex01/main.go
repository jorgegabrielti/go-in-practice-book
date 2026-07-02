package main

import (
	"errors" // Pacote para criar erros
	"fmt"
)

// 1. Função Siples

func boasVindas(nome string) {
	fmt.Printf("Olá, %s! Bem-vindo ao sistema.\n", nome)
}

//2. Múltiplos Retnornos(Cálculo de Retângulo)

// Retorna Área e Perímetro
func geometria(largura, altura float64)(float64, float64) {
	area := largura * altura
	perimetro := 2 * (largura + altura)
	return area, perimetro
}

//3. Tratamento de Erro(Divisão)
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		//erros.New cria um erro simples
		return 0, errors.New("Proibido dividor por zero")
	}
	return a/b, nil
}

//4. Variádica(Média de N notas)
func calcularMedia(notas ...float64) float64 {
	total := 0.0

	if len(notas) == 0 {
		return 0
	}
	for _, nota := range notas {
		total += nota
	}
	return total / float64(len(notas))
}

func main() {
	// Chamada função simpmles(sem retorno)
	boasVindas("Jorge Gabriel")

	// Recebendo múltiplos valores
	a, p := geometria(5.0,3.0) // Go infere que são float64
	fmt.Printf("Retângulo 5x3 -> ÁREA: %.2f, Perímetro: %.2f\n", a, p)

	//Testando erro
	resultado, err := dividir(10,2)

	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("10/2 =", resultado)
	}

	// Forçando erro
	_, err2 := dividir(10, 0) // Usamos _ para ignorar o resultado, pois sabemos que vai falhar.

	if err2 != nil {
		fmt.Println("Tentativa de dividir por 10/0 falhou com sucesso: ", err2)
	}

	// Função variádica
	mediaTurma := calcularMedia(8.5,9.0,5.5,10.0)
	fmt.Printf("Media da turma: %.2f\n", mediaTurma)
}
