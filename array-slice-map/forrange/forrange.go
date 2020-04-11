package main

import "fmt"

func main() {
	numeros := [...]int{9, 2, 3, 4, 5} //compilador conta

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, numero := range numeros {
		fmt.Printf("%d\n", numero)
	}

	for indice := range numeros {
		fmt.Printf("%d\n", indice)
	}
}
