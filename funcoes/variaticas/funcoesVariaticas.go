package main

import "fmt"

//Função Variaticas (Varia a quantidade de parâmetros)
func media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))

}

func imprimirAprovados(aprovados ...string) {
	fmt.Println("\nLista de Aprovados:")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}

}

func main() {

	//Exemplo 1
	resultado := media(7.7, 8.8, 9.9, 0.1)
	fmt.Printf("\nMedia: %.2f", resultado)
	resultado = media(7.7, 8.8, 9.9, 0.1, 12.0, 22.99)
	fmt.Printf("\nMedia: %.2f", resultado)

	//Exemplo 2
	aprovados := []string{"Carlos", "Junior", "Eliana"}
	imprimirAprovados(aprovados...)

}
