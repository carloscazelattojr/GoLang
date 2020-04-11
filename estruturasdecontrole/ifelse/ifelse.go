package main

import (
	"fmt"
	"math/rand"
	"time"
)

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else {
		return "D"
	}
}

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	imprimirResultado(1.3)
	fmt.Println(notaParaConceito(9))

	//If com Init.
	if i := numeroAleatorio(); i > 5 { //também funciona no switch
		fmt.Println("Ganhou", i)
	} else {
		fmt.Println("Perdeu", i)
	}
}
