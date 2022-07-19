package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array 
	s1 := []int{1, 2,3}// slice - parte 
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))


	a2 := [5]int{1, 2, 3, 4, 5}  //criação de um array 
	
	// slice não é um array - SLice define um pedaço de um array

	s2 := a2[1:3]// defindo o slice do indice 2 até 3 
	fmt.Println(a2, s2)

	s3 := a2[:2] //novo slice, mas aponta para mesmo array sempre conta indice 0 , 1 , 2 .....
	fmt.Println(a2, s3)

	// vc pode imaginar um slice como: tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1]
	fmt.Println(s2, s4)

	


}