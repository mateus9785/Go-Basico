package main

import (
	"fmt"
)

func main() {

	//Funcoes
	var numero, outroNumero int = 1, 2
	fmt.Printf("soma de %d + %d = %d \n", numero, outroNumero, somar(numero, outroNumero))
	fmt.Printf("multiplicar de %d * %d = %d \n", numero, outroNumero, multiplicar(numero, outroNumero))
	quociente, resto := dividir(numero, outroNumero)
	fmt.Printf("divisao de %d / %d = %d com resto %d \n", numero, outroNumero, quociente, resto)

	//Closures
	var operacao string = "+"
	var calculo func(n1 int, n2 int) int
	if operacao == "+" {
		calculo = func(n1 int, n2 int) int {
			return n1 + n2
		}
	} else {
		calculo = func(n1 int, n2 int) int {
			return n1 - n2
		}
	}
	fmt.Printf("operacao de %d %s %d = %d \n", numero, operacao, outroNumero, calculo(numero, outroNumero))
}

func buscar() {
	fmt.Println("buscar no banco de dados")
}

func somar(numero int, outroNumero int) int {
	return numero + outroNumero
}

func multiplicar(numero int, outroNumero int) (resultado int) {
	resultado = numero * outroNumero
	return
}

func dividir(numero int, outroNumero int) (int, int) {
	quociente := numero / outroNumero
	resto := numero % outroNumero
	return quociente, resto
}
