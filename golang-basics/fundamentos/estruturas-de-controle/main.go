package main

import (
	"errors"
	"fmt"
)

func main() {

	// If else
	if 2 > 1 {

		fmt.Println("A sintaxe do if não possui paranteses ou algo do tipo para segragar a condição")

	} else {

		fmt.Println("O else deve ser posicionado na mesma linha de fechamenrto do if, caso contrário haverá erro de sintaxe")
	}

	// if com init ()

	if retorno, err := ifinit(); err != nil {

		fmt.Print(retorno)
	}

	// Switch

	lista := []string{"palavra1", "palavra2", "palavra3"}

	switch lista[1] {

	case "palavra1", "palavra2":

		fmt.Println("\nPrimeira condição")
		fallthrough // Muito cuidado ao utilizar essa expressão, pois diferente de outras linguagens, no Go ao cair em um case não haverá continuação para a próxima condição, no entanto
		//  o fallthrough  faz com que siga a execução e vá para o proximo case , porem irá entrar mesmo se não for "verdadeiro".

	case "palavra5":

		fmt.Println("Segunda condição")
		break

	case "palavra3":

		fmt.Println("terceira condição")
	}

	variavel := 5

	switch variavel {

	case 6, 3, 7, 18, 9, 15:
		fmt.Println("primeira condição")

	case 8:

		fmt.Println("segunda condição")

	case 0:

		fmt.Println("terceira condição")

	default:

		fmt.Println("nenhuma das condições acima. O default é semelhante a um else")
	}

}

func ifinit() (string, error) {
	return "ifWithInit", errors.New("erro")
}
