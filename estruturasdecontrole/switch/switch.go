package main

import (
	"fmt"
	"time"
)

func notaParaConceito(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	case bool:
		return "Booleano"
	default:
		return "Não sei o tipo"
	}
}

func main() {
	//Exemplo 1
	fmt.Println(notaParaConceito(10.0))
	fmt.Println(notaParaConceito(8.1))
	fmt.Println(notaParaConceito(1.1))
	fmt.Println(notaParaConceito(5.9))
	fmt.Println(notaParaConceito(11.0))

	//Exemplo 2 - Switch sem associação.
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa Boa noite")
	}

	//Exemplo 3
	fmt.Println(tipo(1.9))
	fmt.Println(tipo("Olá"))
	fmt.Println(tipo(true))
	fmt.Println(tipo(1))
	fmt.Println(tipo(func() {}))
}
