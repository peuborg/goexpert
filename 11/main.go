package main

import "fmt"

type Pessoa struct {
	Id      int
	Nome    string
	Email   string
	Celular string
}

func main() {
	//Structs
	peuborg := Pessoa{
		Id:      1,
		Nome:    "Pedro dos Santos Borges",
		Email:   "email@email.com",
		Celular: "73987654321",
	}

	fmt.Println(peuborg)
}
