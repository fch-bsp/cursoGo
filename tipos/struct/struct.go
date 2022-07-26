package main // Conceito de struct agrupamento de dados

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (recptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() { // Adcionando os produtos classe e produto 
	var escolar produto 
	var eletronico produto
	var bebida produto

	bebida = produto{
		nome: "Redbull",
		preco: 9.80,
		desconto: 0.07,
	}
	escolar = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,	
	}
	eletronico = produto{
		nome: "Noebook",
		preco: 2500.00,
		desconto: 0.10,
	}

	fmt.Println(bebida, bebida.desconto, bebida.preco) // informando a descrição, preço e desconto

	fmt.Println(eletronico, eletronico.precoComDesconto())

	fmt.Println(escolar, escolar.precoComDesconto())

	// ou opção de criar var e declara direto 

	produto2 := produto{"Notebook", 1989.90, 0.10} // declara produto - preço- desconto
	fmt.Println(produto2.precoComDesconto())

	bebida2 := produto{"Agua", 190.00, 00.4}
	fmt.Println(bebida2.precoComDesconto(), bebida2.nome, bebida2.preco)

}
