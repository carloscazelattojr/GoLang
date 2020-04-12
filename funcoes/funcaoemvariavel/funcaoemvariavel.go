package main

import "fmt"

var somar = func(a, b int) int {
	return a + b
}

func main() {

	fmt.Println(somar(4, 5))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(9, 5))

}
