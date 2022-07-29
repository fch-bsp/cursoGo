package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface {} // deixo vazio
	fmt.Println(coisa)

	coisa = 3 // passo um int
	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true //passo bool
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a linguagem de Google"} //passo stuct
	fmt.Println(coisa2)

	




}