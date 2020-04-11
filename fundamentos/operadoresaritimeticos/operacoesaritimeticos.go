package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Somar =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b) //resto

	//bitwise
	fmt.Println("E =>", a&b)   // 11 & 10 == 10
	fmt.Println("Ou =>", a|b)  // 11 | 10 == 11
	fmt.Println("Xou =>", a^b) // 11 ^ 10 == 10

	//outras operações usando math
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))

	c := 3.0
	d := 2.0

	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponencial =>", math.Pow(c, d))

}
