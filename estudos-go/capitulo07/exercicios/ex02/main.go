// O removedor de Itens (Médio - Hacker)
// O Go não tem uma função remove nativa para Slices. Você tem que fazer na mão usando Fatiamento. Crie um slice: nums :=[]int{10, 20, 30, 40, 50}. Imagine que queremos remover o 30 (índice 2). A lógica é: Pegue tudo ANTES do 30 (nums[:2]) e cole com tudo DEPOIS do 30 (nums[3:]). Use append para grudar essas duas metades. Imprima o resultado final. Deve ser [10 20 40 50]

package main

import "fmt"

func main() {
	// Jeito problemático
	//nums := []int{10, 20, 30, 40, 50}
	//fmt.Println("Teste: ", nums[:2])
	//fmt.Println("Teste: ", nums[3:])

	//resultado := append(nums[:2], nums[3:]...)
	//fmt.Println("Resulado: \n", resultado)
	//fmt.Println("Original", nums)

	// Jeito certo
	nums := []int{10, 20, 30, 40, 50}

	resultado := make([]int, 0, len(nums)-1)

	resultado = append(resultado, nums[:2]...)
	resultado = append(resultado, nums[3:]...)

	fmt.Println("Resultado: ", resultado)
}
