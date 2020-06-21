package main

import (
	estrutura "carros"
	"fmt"
)

func main() {
	carro := estrutura.Carro{Modelo: "Palio", Marca: "Fiat"}
	opcao := 0
	for opcao != 3 {
		fmt.Println("O que deseja fazer ?")
		fmt.Println(" 1 - Acelerar")
		fmt.Println(" 2 - Frear")
		fmt.Println(" 3 - Sair")
		fmt.Scanf("%d", &opcao)
		var err error = nil
		switch opcao {
		case 1:
			err = fazerAcelerar(&carro)
		case 2:
			err = fazerFrear(&carro)
		}
		if err != nil {
			fmt.Printf(" ** Não foi possível executar a ação %s ** \n", err)
		} else if opcao != 3 {
			fmt.Printf("O carro %s da marca %s está a %.2f km/h \n", carro.Modelo, carro.Marca, carro.Velocidade)
		}
	}
	fmt.Println("Fim da execução")
}

func fazerAcelerar(veiculo estrutura.Acelerador) error {
	return veiculo.Acelerar()
}

func fazerFrear(veiculo estrutura.Freador) error {
	return veiculo.Frear()
}
