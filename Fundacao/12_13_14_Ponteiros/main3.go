package main

import "fmt"

type Pessoa struct {
	Id      int
	Nome    string
	Email   string
	Celular string
}

func (p *Pessoa) resetarCelular(cel string) {
	p.Celular = cel
	fmt.Println(p)
}

func NewPessoa() *Pessoa {
	return &Pessoa{Id: 0,
		Nome:    "zerado",
		Email:   "null@email",
		Celular: "99999999999",
	}
}

func main() {
	//Structs
	peuborg := Pessoa{
		Id:      1,
		Nome:    "Pedro dos Santos Borges",
		Email:   "email@email.com",
		Celular: "73987654321",
	}
	peuborg.resetarCelular("123456")
	fmt.Println(peuborg)

	novato := NewPessoa()
	fmt.Println(novato)
}
