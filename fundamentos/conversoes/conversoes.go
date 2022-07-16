package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
    y := 2
	fmt.Println(x / float64(y))

    nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//Cuidado....
	fmt.Println("Teste " + string(97))

	//int para String
    fmt.Println("Teste " + strconv.Itoa(123))



	//string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 123)

	b, _ := strconv.ParseBool("true")//ou valor 1-0 true-false
	if b {
		fmt.Println("Verdadeiro")
	}









}

