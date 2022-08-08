package main

import (
	"fmt"
	"time"
)

// Channel (Canal) - é a forma de comunicação entre goroutines
// Channel é um tipo

func doisTResQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)  // espera por 1 segundo 
	c <- 2 * base //enviando para o channel 

	time.Sleep(time.Second)  // espera por 1 segundo 
	c <- 3 * base  //enviando para o channel

	time.Sleep(time.Second)  // espera por 1 segundo 
	c <- 4 * base  //enviando para o channel
}



func main() {
	c := make(chan int)
	go doisTResQuatroVezes(2, c)
	fmt.Println("Primeira etapa")

	a, b := <-c, <-c // criando var e recebendo os dados do channel
	fmt.Println("Segunda etapa") 
	fmt.Println(a)
	fmt.Println("Terceira etapa")
	fmt.Println(b)

	fmt.Println(<-c) //rebendo os dados do channel
	fmt.Println("Quarta etapa")
}