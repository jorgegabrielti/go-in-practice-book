package main

import "fmt"

var login string = "admin"
var senha string = "12345"

func main() {
	if login == "admin" && senha == "12345" {
		fmt.Println("Acesso liberado!")
	} else {
		fmt.Println("Login ou senha inválidos.")
	}
}
