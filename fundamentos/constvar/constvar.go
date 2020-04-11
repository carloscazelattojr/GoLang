package main

import (
	"fmt"
	"math"
)

//Ao declarar uma variavel/constante, é obrigatório utilizar, senão gera erro de compilação
func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo (float64) inferido pelo compilador

	//forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2)
	fmt.Println("Area  da circunferencia é", area)

	//Formas de declaracao - 2
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	//Formas de declaracao - 3
	var e, f bool = true, false
	fmt.Println(e, f)

	//Formas de declaracao - 4
	g, h, i := 2, false, "epa"
	fmt.Println(g, h, i)

}
