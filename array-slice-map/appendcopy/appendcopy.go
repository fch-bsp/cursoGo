package main //defindo o pagote main

import "fmt"

func main() { //defindo funão mais para ter um código executável
	array1 := [3]int{1, 2, 3} //criando uma array de 3 posição tamnho fixo
	var slice1 []int

	//array 1 = append(array1, 4, 5, 6)

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2)

}
