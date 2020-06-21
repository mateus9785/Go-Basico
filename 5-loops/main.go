package main

import "fmt"

func main() {
	var numero int = 6
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", i, numero, numero*i)
	}

	//WHILE
	j := 0
	for j <= 10 {
		fmt.Printf("%d x %d = %d \n", j, numero, numero*j)
		j++
	}
}
