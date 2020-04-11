package main

import "fmt"

func main() {
	//var aprovados map[int]string

	aprovados := make(map[int]string)

	//Maps devem ser inicializados
	aprovados[12345678978] = "Maria"
	aprovados[10002378978] = "Pedro"
	aprovados[12222078978] = "Joao"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[10002378978])
	delete(aprovados, 10002378978)
	fmt.Println("Deletado: " + aprovados[10002378978])

}
