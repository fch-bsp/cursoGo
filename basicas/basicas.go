package main

import (
	"fmt"
)

// tipos de funções

func F1() {
	fmt.Print("F1") // ""um bloco"" Não recebe nada como estrada de dados e não retorna nada
}

func F2(p1 string, p2 string) { // ""um bloco"" Recebe parametros e não retorna nada
	fmt.Printf("F2: %s %s \n", p1, p2)

}

func f3() string { // ""um bloco""Não recebe parametros e recebe alguma coisa
	return "F3"
}

func f4(p1, p2 string) string { // ""um bloco"" Recebe dois parametros e retorna um único parametro
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) { //Não recebe nenhum parametro e retorna dois valores tipo string
	return "Retorno1", "Retrono2" //chamando o return precisa passar dois valores

	r3, r4 := f3(), f4("Param1", "Param2")
	println(r3)
	println(r4)

}
