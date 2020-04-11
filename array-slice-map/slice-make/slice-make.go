package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	//Append para completar.
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	//Append para aumentar. Por padrão ele dobra de tamanho
	s = append(s)
	fmt.Println(s, len(s), cap(s))

	//referenciando Slice x Array
	s1 := make([]int, 10, 20)
	s2 := append(s1, 2, 3, 4)
	fmt.Println(s1, s2)

	s1[0] = 7
	fmt.Println(s1, s2)

}
