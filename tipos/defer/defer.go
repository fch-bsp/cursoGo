package main

import "fmt"

// Resumo: adicionar antes de uma sentença de código atrasar

func obterValorAprovado(numero int) int { // criação de fun var numer int retorna int
	defer fmt.Println("Fim!!") // aguarda o laço
	if numero >= 5000 {         // se numer for maior que 5000 faz isso
		fmt.Println("Valor alto....")
		return 5000  // retorna o valor 

	} else { // caso contrario faz isso
		fmt.Println("Valor baixo...")
		return numero  //retirna a var numero 

	}

}	

func main() {
	fmt.Println(obterValorAprovado(5020))
	fmt.Println(obterValorAprovado(3000))





}
