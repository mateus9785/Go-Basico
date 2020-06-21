package main

import (
	"errors"
	"fmt"
)

type Acelerador interface {
	acelerar() error
}

type Freador interface {
	frear() error
}

type Carro struct {
	modelo     string
	marca      string
	velocidade float32
}

func (carro *Carro) acelerar() error {
	if carro.velocidade < 15 {
		carro.velocidade += 5
		return nil
	} else {
		return errors.New("o carro já está em sua velocidade máxima")
	}
}

func (carro *Carro) frear() error {
	if carro.velocidade > 0 {
		carro.velocidade -= 5
		return nil
	} else {
		return errors.New("o carro está parado")
	}
}

func main() {
	carro := Carro{modelo: "Palio", marca: "Fiat"}
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
			fmt.Printf("O carro %s da marca %s está a %.2f km/h \n", carro.modelo, carro.marca, carro.velocidade)
		}
	}
	fmt.Println("Fim da execução")
}

func fazerAcelerar(veiculo Acelerador) error {
	return veiculo.acelerar()
}

func fazerFrear(veiculo Freador) error {
	return veiculo.frear()
}
