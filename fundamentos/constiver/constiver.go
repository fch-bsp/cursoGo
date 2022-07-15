package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.145
	var raio = 3.2 // tipo (float64) inferido pelo combilador

	// forma reduzida de criar uma var

	area := PI * m.Pow(raio, 2)                  // := voce declara e atribuir var
	fmt.Print("A área de cicunferência é", area) //quando define uma var precisa utiliza-la

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false//declarei a var e o tipo inicializei os valores 
	fmt.Println(e, f)

	g, h, i := 2, false, "apa!"//declarei 3 tipo de var tipos: int- bool-string
	fmt.Println(g, h, i)

}
