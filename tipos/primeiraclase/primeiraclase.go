package main

import "fmt"

var soma = func(a, b int) int { //criei uma var nome de soma - retorno é "a + b"
	return a + b

}

func main() {
	fmt.Println(soma(2, 3))

	sub := func(a, b int) int {  //executando uma func main dentro de uma função local
		return a - b

	}

	fmt.Println(sub(8, 5))
}
