package main

import (
	"errors"
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {

		fmt.Println(i)
	}

	// while  - não existe, sendo assim podemos fazer um while adaptado

	while := 1

	for while <= 10 {

		fmt.Println(while)
		while++

	}

	// do-while

	anExpression := false

	for ok := true; ok; ok = anExpression {

		fmt.Println("entrou no loop")
	}

	// for com range

	variavel2 := map[string]string{
		"palavra1": "jade",
		"palavra2": "do pai dela",
		"palavra3": "maravilhosa",
	}

	for chave, valor := range variavel2 {

		fmt.Println(chave, valor)

	}

	valor, erro := Teste2()

	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println(valor)
	}

	// resultado := soma(2, 3)

	// fmt.Println("O resultado da soma é:", resultado)

	imprimirHello := imprimirMensagem("Olá, mundo!")
	// Chamando a função retornada para imprimir a mensagem
	imprimirHello()

}

func teste() (string, error) {
	return "teste", errors.New("erro de sintaxe na função teste")
}

//  retorno nomeado

func Teste2() (retornoString string, retornoErro error) {
	retornoString = "teste2"
	retornoErro = nil

	return
}

// função com parâmetros

func soma(a int, b int) int {
	return a + b
}

func imprimirMensagem(mensagem string) func() {
	return func() {
		fmt.Println(mensagem)
	}
}
