package main

import (
	"fmt"

	"github.com/peuborg/goexpert/06_Packaging/1/math"

	"github.com/google/uuid"
)

func main() {
	//Escopo público
	println("Alo mundo - público")
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())

	//Escopo privado
	println("Alo mundo - privado")
	m1 := math.NewMath(1, 2)
	fmt.Println(m1)

	//
	println(uuid.New().String())
}
