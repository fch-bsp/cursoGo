package main

import "fmt"

func main(){
	numeros := [...]int{1, 2, 3, 4, 5} //compilador conta!

	for i, numero := range numeros { //acessando o indice do "i" e numeros 
		fmt.Printf("%d) %d\n", i+1, numero)
		
	}
     // '_ ' para ignorar sem precisar usar indice 
	for _, num := range numeros { // trabalhar apenas com numeros e não com indice 
		fmt.Println(num)
	}

/*  "_" serve caso precisa declara uma variavél e não usar por padrão o Go quando você declara uma variavél 
você precisa usar !!
 */



}