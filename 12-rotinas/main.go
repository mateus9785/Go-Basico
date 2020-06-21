package main

import (
	"fmt"
)

func main() {
	var limite int
	canal1 := make(chan int, 2000) // buffer que é indica quantas vezes pode receber valor
	canal2 := make(chan int, 500)  // pois canal é sincrono e trava execucao das rotinas
	fmt.Print("Informe um limite")
	fmt.Scanf("%d", &limite)
	go conteAte(limite, canal1)
	go func(n int) {
		resultado := 0
		for i := 0; i <= n*10; i++ {
			fmt.Printf(" - [anonimo] O número é %d \n", i)
			resultado = i * 10
		}
		canal2 <- resultado
	}(limite)
	for i := 0; i <= limite*10; i++ {
		fmt.Printf(" - [main] O número é %d \n", i)
	}
	fmt.Printf("Canal 1 = %d, Canal 2 = %d \n", <-canal1, <-canal2)
}

func conteAte(limite int, canal chan int) {
	for i := 0; i <= limite*20; i++ {
		fmt.Printf(" - [conteAte] O número é %d \n", i)
	}
}
