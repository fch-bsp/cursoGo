package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("String:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 3)
	fmt.Println(">=", 3 >= 2 )

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)

	type Pessoa struct {
		nome string
	}
	
	p1 := Pessoa{"Fernando"}
	p2 := Pessoa{"Fernando"}
	fmt.Println("Pessoas:", p1 == p2)

}