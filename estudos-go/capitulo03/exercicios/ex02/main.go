// Você está medindo chuva.
// - Ontem choveu 10.5mm.
// - Hoje choveu 10mm(inteiro). Crie variáveis para essea valores. tente somá-las em uma variável total sem fazer conversões. Veja o erro. Corrija o código fazendo o "casting" necessário para procesar a soma como float64. Imprima o total.
package main

import "fmt"

var chuvaOntem float64 = 10.5
var chuvaHoje int = 10
var total float64 = chuvaOntem + float64(chuvaHoje) // Casting => float64(VALOR)

func main(){
	fmt.Println("Resultados das chuvas de ontem + hoje: ", total)
}