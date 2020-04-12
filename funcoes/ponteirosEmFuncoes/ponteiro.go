package main

import "fmt"

func inc1(n int) {
	n++
}

//Ponteiro é representado por um *
//* é utilizado para desreferenciar, ter acesso ao valor dem vez do endereco da memórioa
func inc2(n *int) {
	*n++
}

func main() {

	n := 1
	inc1(n)
	fmt.Println(n)

	//quando eu utilizo o & é passado o endereco da variavel
	inc2(&n)
	fmt.Println(n)

}
