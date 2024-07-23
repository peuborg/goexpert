package main

import (
	"18_Pacotes/matematica"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Volkswagen"}
	carro.Marca = "T-Cross"

	fmt.Printf("Resultado: %v", s)
	println(matematica.A)
	fmt.Println(carro)
	carro.Andar()
	fmt.Println(uuid.New())
}

/*
comandos terminal:
	go run main.go
	go mod init 18_Pacotes/matematica
	go mod tidy		// Atualiza os requires necessários para a aplicação go
	go get golang.org/x/exp/constraints
	go get github.com/google/uuid v1.6.0
*/
