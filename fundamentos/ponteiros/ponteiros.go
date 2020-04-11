package main

import "fmt"

func main() {
	i := 1

	//go não tem aritimética de ponteiros.
	var p *int = nil

	p = &i //pegando o endereco da variavel i

	*p++ //Desreferenciando (Pegando o valor e adicionando)
	i++

	fmt.Println(p, *p, i, &i)

}
