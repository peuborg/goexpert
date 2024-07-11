package main

import "fmt"

type Cliente struct {
	Codigo        int
	DadosPessoais Pessoa
}

type Pessoa struct {
	Id      int
	Nome    string
	Email   string
	Celular string
}

func main() {
	//Composição de Structs
	peuborg := Pessoa{
		Id:      1,
		Nome:    "Pedro dos Santos Borges",
		Email:   "email@email.com",
		Celular: "73987654321",
	}

	cli := Cliente{
		Codigo:        1,
		DadosPessoais: peuborg,
	}

	fmt.Println(cli)
}
