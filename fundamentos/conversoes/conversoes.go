package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2

	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado
	fmt.Println("Teste" + string(123))

	//Converter Inteiro para String
	fmt.Println("Teste " + strconv.Itoa(123))

	//Converter String para Inteiro
	num, _ := strconv.Atoi("123") //Atenção, ao converter String para Inteiro, ele retorna 2 valores. O primeiro é o valor e o segundo é para erro(Se houver na conversao).
	fmt.Println(num - 1)

	b, _ := strconv.ParseBool("true") //Idem, retorna dois retorno.
	if b {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("Falso")
	}

}
