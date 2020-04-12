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

	//-------------------Exemplo 2-----------------------------

	funcsESalarios := map[string]float64{
		"Jose":    1125.45,
		"Garbiel": 2251.99,
		"Pedro":   1200.00,
	}
	fmt.Println(funcsESalarios["Jose"])
	funcsESalarios["Jose"] = 5000.00

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

	//-------------------Exemplo 3-----------------------------
	//---- MAP Aninhado

	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel": 15456.78,
			"Guga":    8452.55,
		},
		"J": {
			"Junior": 5567.60,
		},
		"C": {
			"Carlos": 22654.22,
		},
	}

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
