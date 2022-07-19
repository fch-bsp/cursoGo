package main

import "fmt"

func main() { //metodo de maps com chave e valor 
	funcsPorletra := map[string]map[string]float64{ //criei um mapa "F" outro map "Nome" string e salario float64 
		"F": {
			"Fenando Horas": 15852.78,
		},
		"P": {
			"Priscila Santos": 1254.41,
		},
		"K": {
			"Klara dos Santos": 214587.2584,
		},
	}

	delete(funcsPorletra, "P")

	for letra, funcs := range funcsPorletra {
		fmt.Println(letra, funcs)
	}
}