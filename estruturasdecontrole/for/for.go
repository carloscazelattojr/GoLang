package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { //não tem while em GO, mas o for se comporta como se fosse o While
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	//laço infinito
	for {
		fmt.Println("Imprimindo infinitamente...")
		time.Sleep(time.Second)
	}

}
