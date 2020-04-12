package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}

	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, qtde: 2, preco: 23.49},
			item{produtoID: 2, qtde: 8, preco: 229.77},
			item{produtoID: 4, qtde: 22, preco: 3.09},
		},
	}

	fmt.Printf("Valor total é R$ %.2f", pedido.valorTotal())
}
