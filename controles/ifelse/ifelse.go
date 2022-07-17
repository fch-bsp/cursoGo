package main // Chamando o pagote main

import "fmt" // Importanto o "fmt" - formate

func imprimirResultado(nota float64) { //criando uma função e informando var 
	if nota >= 7 { // se a nota for maior que 7 entra nesse laço 
		fmt.Println("Aprovado!! com a nota", nota)
	} else { // caso contrario entra nesse laço
		fmt.Println("Reprovado!! com a nota", nota)
	}
}

func main() {// precisa criar a função MAIN 
	imprimirResultado(7.1)// nota final 
	imprimirResultado(5.1)
}
