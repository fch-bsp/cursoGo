package main

import "fmt"

type item struct { // definindo a struct
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct { // definindo a struct
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens { //descarta indice "_" e traz apenas o produto
		total += item.preco * float64(item.qtde)
	}

	return total

}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID:1, qtde:2, preco:12.10}, // adcionando item/qtde/preço
			item{produtoID:2, qtde:1, preco:23.49},
			item{produtoID:3, qtde:100, preco:3.13},
		},
	}

	fmt.Printf("Valor total de pedido é R$ %.2f", pedido.valorTotal())


}
