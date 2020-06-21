package main

import (
	"fmt"
)

func main() {
	// grupo de inteiros SEM SINAL
	var u1 uint8 = 255 //alias byte
	fmt.Println(u1)
	var u2 uint16 = 33
	fmt.Println(u2)
	var u3 uint32 = 22
	fmt.Println(u3)
	var u4 uint64 = 11
	fmt.Println(u4)

	// grupo de inteiros COM SINAL
	var u5 int8 = 127
	fmt.Println(u5)
	var u6 int16 = 33
	fmt.Println(u6)
	var u7 int32 = 22 //rune
	fmt.Println(u7)
	var u8 int64 = 11
	fmt.Println(u8)

	//Depende da arquitetura do PC
	var t1 uint = 12
	fmt.Println(t1)
	var t2 int = 24
	fmt.Println(t2)

	//flutuantes
	var f1 float32 = 2
	fmt.Println(f1)
	var f2 float64 = 2
	fmt.Println(f2)
	var c1 complex64 = 2
	fmt.Println(c1)
	var c2 complex128 = 2
	fmt.Println(c2)

	//string
	var s1 string = "Teste"
	fmt.Println(s1[0]) //84 pois responde codigo ASC

	//Boleanos
	var b1 bool = true
	fmt.Println(b1)

	//var nome, sobrenome string = "mateus", "oliveira"
	var (
		nome      string = "Mateus"
		sobrenome string = "Oliveira"
	)

	fmt.Println(nome + " " + sobrenome)

	const constante = 35

	fmt.Println(constante)

	//inferencia
	var nome = "teste"
	fmt.Println(nome)
	nomeCompleto := "mateus oliveira"
	fmt.Println(nomeCompleto)
}
