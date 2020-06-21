package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		numero1 int
		numero2 int
	)
	fmt.Print("Digite o primeiro número : ")
	fmt.Scanln(&numero1)
	fmt.Print("Digite o segundo número : ")
	fmt.Scanln(&numero2)
	fmt.Printf("%d + %d = %d \n", numero1, numero2, numero1+numero2)
	fmt.Printf("%d - %d = %d \n", numero1, numero2, numero1-numero2)
	fmt.Printf("%d * %d = %d \n", numero1, numero2, numero1*numero2)
	fmt.Printf("%d / %d = %d \n", numero1, numero2, numero1/numero2)
	fmt.Printf("%d %% %d = %d \n", numero1, numero2, numero1%numero2)

	numero3 := 5
	numero3 += numero1
	numero3 -= numero2

	var texto string
	fmt.Print("Digite um número : ")
	fmt.Scanln(&texto)
	// var numero int
	// numero, _ = strconv.Atoi(texto)
	numero, _ := strconv.ParseInt(texto, 10, 32)
	var conv = float64(numero)
	fmt.Println(conv)

}
