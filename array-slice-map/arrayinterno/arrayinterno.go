package main

import "fmt"

func main() {
	
	//criar um slice com do tipo inteiro com 10 posições e um array com 20 posições
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))


	s1[0] = 7
	fmt.Println(s1, s2)




}