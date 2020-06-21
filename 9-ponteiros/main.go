package main

import (
	"fmt"
)

func main() {
	// var saldo float64
	var saldo = new(float64)
	fmt.Print("Digite o seu saldo: ")
	//fmt.Scanf("%f", &saldo)
	fmt.Scanf("%f", saldo)
	//calculaRendimento(&saldo);
	calculaRendimento(saldo)
	fmt.Printf("Seu saldo com rendimento é de %.2f \n", *saldo)
}

func calculaRendimento(saldo *float64) {
	*saldo += *saldo * 0.03
}
