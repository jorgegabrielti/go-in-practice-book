// Crie um variável do tipo uint8 para representar a quilometragem do seu carro (sim, é um carro com hodômetro muito pequeno, que só vai até 255km). Comece com o valor 250. Adicione 10 à variável (quilometragem = quilometragem + 10). Imprima o resultado. Explique nos comentários do seu código por que o resultado foi estranho.
package main

import "fmt"

var quilometragem uint8 = 250

func main() {

	fmt.Println("Kilometragem =", quilometragem + 10) // 250 + 10 = 260 - 256 = 4. Por conta da limitação de tamanho do uint8, o contador vira e iniciar a contagem do 0.
}