package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool { // recebe valor inteiro e transforma em um booleano
	for i := 2; i < num; i++ { //recebe num se esse numero for divissivél false  por 2 se não true
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {  // criação de uma laço 
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Microsecond * 180)
				break
			}

		}
	}
	close(c)  // quando deixa de usar um canal fechar o mesmo caso contário pode gerar um detlook 
}


func main() {
	c := make(chan int, 30)
	go primos(cap(c), c)
	for primo := range c {  // criação de um laço for em cima de um canal 
		fmt.Printf("%d", primo)
	}
	fmt.Println("  Fim!")
}