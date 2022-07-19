package main

import "fmt"



func main() {
	funcESalarios := map[string]float64{
		"Jose João": 1125.25,
		"Gabriela SIlva": 2125.45,
		"Pedro Junior": 1325.25,


	}

	funcESalarios["Rafael Filho"] = 1350.0
	delete(funcESalarios, "Inexistente")

	for nome, salario := range funcESalarios { //criação de for pergorrendo nome e salario
		fmt.Println(nome, salario)
	}

	for _, salario := range funcESalarios { //criação de for percorrendo apenas o salario
		fmt.Println(salario)
	}

	for nome := range funcESalarios { //criação de for percorrendo apenas o nome
		fmt.Println(nome)
	}



}