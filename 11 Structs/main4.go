package main

import "fmt"

type Pessoa struct {
	Id      int
	Nome    string
	Email   string
	Celular string
}

type ColetorLixo interface {
	Resetar()
}

func (p Pessoa) Resetar() {
	p.Id = 0
	p.Nome = ""
	p.Email = ""
	p.Celular = ""
	fmt.Println(p)
}

func Limpeza(limpador ColetorLixo) {
	limpador.Resetar()
}

func main() {
	//Interfaces em Structs
	peuborg := Pessoa{
		Id:      1,
		Nome:    "Pedro dos Santos Borges",
		Email:   "email@email.com",
		Celular: "73987654321",
	}

	Limpeza(peuborg)

	fmt.Println(peuborg)
}
