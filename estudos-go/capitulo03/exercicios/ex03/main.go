// O vergo %T do pacote fmt serve para imprimir o TIPO de uma variável. Crie um programa que declare. 
// x := 10
// y := 10.0
// z := "10" Imprima o valor e o tipo de cada usando fmt.Printf("Valor: %v, Tipo: %T\n", variavel, variavel). 
// Observe como a inferência de tipos do Go decide se algo é inte ou float64 automaticamente.

package main

import "fmt"

func main(){
  x := 10
  y := 10.0
  z := "10"

  fmt.Printf("Valor: %v, Tipo: %T\n", x, y, z)
}