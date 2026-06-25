package main

import "fmt"

var total = 10

func main() {
  total := 5
  fmt.Println(total) //Imprime 5 e não 10.
  
  imprimeTotal()
}

func imprimeTotal(){
  fmt.Println(total) //Imprime 10 e não 5.
}