package main

type produto struct {
	nome string
	preco float64
	desconto float64
}

// Método: função com receiver (recptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}	

func main() {
	proproduto1 = prproduto
	  nome:  "Lapis",
	  preco:  1.79,
	  desconto: 0.05,

    }

	fmt.Println(proproduto1)

}