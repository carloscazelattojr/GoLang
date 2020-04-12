package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcaoNomeado func(int, int) int, p1, p2 int) int {
	return funcaoNomeado(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 3, 4)
	fmt.Println(resultado)
}
