package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas deve ser inicializados

	aprovados := make(map[int]string)

	aprovados[25125854485] = "Fernando"
	aprovados[25541154485] = "Priscila"
	aprovados[52525854485] = "Klara"


	for cpf, nome := range aprovados { //criando uma var nome "cpf" e "nome"
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
	fmt.Println(aprovados[25541154485])
	delete(aprovados, 25541154485)
	fmt.Println(aprovados[25541154485])
}