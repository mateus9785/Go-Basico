package main

import "fmt"

func main() {
	var numero uint8 = 5
	if numero < 0 {
		fmt.Println("menor que zero")
	} else if numero > 0 && numero < 9 {
		fmt.Println("maior que zero e menor que 9")
	} else {
		fmt.Println("outros casos")
	}

	switch numero {
	case 1:
		fmt.Println("igual a 1")
	case 2, 3, 4:
		fmt.Println("outros casos")
	default:
		fmt.Println("inválido")
	}
}
