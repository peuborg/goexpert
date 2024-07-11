package main

import "fmt"

type Pessoa struct {
	Id      int
	Nome    string
	Email   string
	Celular string
}

func (p Pessoa) Resetar() Pessoa {
	p.Id = 0
	p.Nome = ""
	p.Email = ""
	p.Celular = ""
	fmt.Println(p)

	return p
}

func main() {
	//Métodos em Structs
	peuborg := Pessoa{
		Id:      1,
		Nome:    "Pedro dos Santos Borges",
		Email:   "email@email.com",
		Celular: "73987654321",
	}

	fmt.Println(peuborg)

	peuborg = peuborg.Resetar()

	fmt.Println(peuborg)
}
