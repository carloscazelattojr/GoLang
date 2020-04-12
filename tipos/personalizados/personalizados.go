//meu tipo personalizado.package personalizados

package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.99, 5.0) {
		return "b"
	} else {
		return "E"
	}
}

func main() {

	fmt.Println(notaParaConceito(9.8))
}
