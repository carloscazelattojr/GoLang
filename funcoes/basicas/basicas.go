package main

import "fmt"

//Basuca
func f1() {
	fmt.Println("Imprimindo F1")
}

//Com parametros
func f2(p1 string, p2 string, p3 int) {
	fmt.Println("\nImprimindo F2: ", p1, p2, p3)
}

//Com retorno
func f3() string {
	return "\nImprimindo F3"
}

//Parametros e com retorno
func f4(p1, p2 string) string {
	fmt.Printf("F4: %s %s", p1, p2)
	return "F4 OK\n"
}

//COm dois retornos
func f5() (string, string) {
	return "\nF5: Retorno 1", "Retorno 2"
}

func main() {
	f1()

	f2("Carlos", "Junior", 36)

	r3 := f3()
	fmt.Println(r3)
	fmt.Println(f3())

	f4("Carlos", "Junior")

	r5, r6 := f5()
	fmt.Println(r5, r6)
	_, r2 := f5() //Desconsidero a primeira variavel de retorno.
	fmt.Println(r2)
	r5a, _ := f5() //Desconsidero a primeira variavel de retorno.
	fmt.Println(r5a)

	fmt.Println(f5())

}
