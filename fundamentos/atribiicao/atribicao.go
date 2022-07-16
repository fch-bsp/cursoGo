package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Print(b)

	i := 3 //inferência de tipo
	i += 3 // i = i + 3 atribuição mesmo que i recebe i + 3
	i -= 3 // i = i - 3 atribuição mesmo que i recebe i - 3
	i /= 2 // i = i / 2 atribuição mesmo que i recebe i / 2
	i *= 2 // i = i * 2 atribuição mesmo que i recebe i * 2
	i %= 2 // i = i % 2 atribuição mesmo que i recebe i % 2

    fmt.Println(i)

    x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x 
	fmt.Println(x, y)
}