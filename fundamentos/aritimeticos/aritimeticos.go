package main

import "fmt"

func main() {
    a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multplicação =>", a*b)
	fmt.Println("Módulo=>", a%b)


	//bitwise
    fmt.Println("E =>", a&b)  // 11 & 10 = 10
	fmt.Println("Ou =>", a|b) // 11 | 10 = 11
	fmt.Println("Xor =>",a^b) // 11 ^ 10 = 01

    c := 3.0
	d := 2.0

	//outras operações usando o match...

	fmt.Println("Maior =>", match.Max(float64(a), float32(b)))
	fmt.Println("Menor =>", match.Min(c , d))
    fmt.Println("Exponenciação =>", match.Pow(c, d)) 
	
}