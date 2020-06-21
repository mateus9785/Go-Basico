package main

import (
	"fmt"
)

func main() {

	//Defer - executado no final
	fmt.Println("abrir conexão com banco de dados")
	defer fmt.Println("fechar conexão com banco de dados")
	buscar()

	const condicao = 10
	//Panic
	if condicao > 10 {
		panic("Numero maior que 10")
	}

	//Recover para tratar exceptions

	defer func() {
		if ex := recover(); ex != nil {
			fmt.Printf("ocorreu um erro : %s", ex)
		}
	}()
	if condicao > 10 {
		panic("Numero maior que 10")
	} else {
		fmt.Println("bom número")
	}
}

func buscar() {
	fmt.Println("buscar no banco de dados")
}
