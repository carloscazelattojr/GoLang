package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//nímeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos) ... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal... int8, int16, int36, int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela Unicode(int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	//números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo do x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	//valores boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)
	var bo2 bool = false
	fmt.Println(bo2)

	//string
	s1 := "Olá meu nome é Carlitos"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	//String com multiplas linhas
	s2 := `Olá
	Meu 
	nome é
	Carlitos`
	fmt.Println("O tamanho da string é", len(s2))
	fmt.Println(s2)

	var s3 string = "Carlos"
	fmt.Println(s3)

	//Char....
	char := 'a' // Se eu utilizar Aspa simples '', devo colocar somente 1 caracter, aspas somente com 1 caracter, se for mais, tem que colocar aspas duplas ""
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)

}
