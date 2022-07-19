package main

import (
	"fmt"
)

func main() { // " := " cria uma variavél, difine e declara

	s := make([]int, 10) // cria um slice com 10 elementos só que o array interno é o numero 12
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s)) //imprindo o "len" tamanho "cap" capacidade

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

}
