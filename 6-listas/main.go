package main

import (
	"container/list"
	"fmt"
)

func main() {
	//Vetor
	numeros := [5]int{10, 20, 30, 40, 50}

	for _, numero := range numeros {
		fmt.Printf("%d \n", numero)
	}

	fmt.Printf("tamanho de matriz %d", len(numeros))

	//Slice

	amigos := make([]string, 3)
	amigos[0] = "Paulo"
	amigos[1] = "Júlio"
	amigos[2] = "camila"
	amigos = append(amigos, "Carlos")
	amigos = append(amigos, "Larissa", "Ana")

	outroSlice := amigos[2:4]
	fmt.Println(outroSlice)

	maisUmSlice := amigos[:4]
	fmt.Println(maisUmSlice)

	//Mapa
	mapacurso := make(map[string]int)
	mapacurso["teste"] = 1

	//Lista
	lista := list.New()
	elemento := lista.PushBack(1)
	lista.PushFront(0)
	lista.PushBack(2)
	lista.Remove(elemento)
	for el := lista.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}

}
