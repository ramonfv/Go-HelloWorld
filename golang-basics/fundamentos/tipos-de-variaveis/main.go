package main

import (
	"fmt"
)

func main() {

	//  tipos de variáveis
	// var test float32 = 3659.9632
	// var inteiro int16 = 25
	// var palavras String = "Hello World"

	// instanciando variáveis corretamente ou usualmente

	// newVar := "assim"
	// newVar := 59
	newVar := map[string]string{

		"teste": "teste",
	}

	fmt.Printf("%T", newVar)

	// Tipo any
	var qualquerValor interface{}

	qualquerValor = 20.5

	fmt.Print("\n", qualquerValor)

	// Arrays e slices

	var lista [4]string = [4]string{"string1", "string2", "string3", "string4"}

	fmt.Println("\nTamanho:", len(lista))
	fmt.Println("Capacidade suportada:", cap(lista))

	var lista2 []int = []int{1, 2, 3, 4}

	fmt.Println("\nTamanho:", len(lista2))
	fmt.Println("Capacidade suportada:", cap(lista2))

	lista2 = append(lista2, 5)

	fmt.Println("\nTamanho:", len(lista2))
	fmt.Println("Capacidade suportada:", cap(lista2))

	// um array declarado sem um tamanho definido é escalável em recurso, ao estourar o limite de memória a quantidade inicialmente alocada  é dobrada

	// Structs

	var pessoa Pessoa = Pessoa{

		nome:     "Jade",
		idade:    5,
		cpf:      "111.111.111.11",
		ocupacao: "herdeira",
	}

	fmt.Print(pessoa)

}

// Struct

type Pessoa struct {
	nome     string
	idade    int
	cpf      string
	ocupacao string
}

type User struct {
	Email    string
	password string
}

// Variáveis públicas e privadas

func privada() {
	fmt.Println("Para uma função ser privada, a sua declaração deve possuir a primeira letra minúscula.")
}

func Publica() {
	fmt.Println("Para uma função ser publica, a sua declaração deve possuir a primeira letra maiúscula.")
}
