package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numeros insteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é:", reflect.TypeOf(32000))

	// sem sinal (só positivo).... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int 8 int 16 int32 int64

	i1 := math.MaxInt64

	fmt.Println("O valor máximo de int é", i1)
	fmt.Println("O valor de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela UNicode (int32)
	fmt.Println("Prune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo literal de 49.99 é", reflect.TypeOf(49.99)) //float64

	// Boolean
	bo := false
	fmt.Println("O valor de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	s1 := "Olá meu nome é Fernando Horas"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))
    

	nome := "Meu nome é Fernando"
	fmt.Println(nome + "!")



	//String com multiplas linhas

	s2 := ` OLá
	meu 
	nome
	é 
	Fernando
	Horas
	
	`

    fmt.Println("O tamanho da string é", len(s2))
	fmt.Println(s2 + "!!")

	





}
