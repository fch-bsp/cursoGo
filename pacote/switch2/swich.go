package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now() // peguei a hora atual a armazenei na var "t"
	switch { 
	case t.Hour() < 12: // Se a hora for menor que 12 
		fmt.Println("Bom dia!!")//informa esse resultado
	case t.Hour() < 18: // Se a hora for menor que 18
		fmt.Println("Boa Tarde!!")//informa esse resultado
	default: // Se a hora não for nenhum dos dois laços de cima 
		fmt.Println("Boa Noite!!")//informa esse resultado			
	}


}


