package main

import "fmt"

func obterResusultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"

	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResusultado(6))
}
