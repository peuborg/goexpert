package main

import "fmt"

var (
	conectou bool    = true
	tempo    int     = 10
	resposta string  = "sucesso"
	lucro    float64 = 500.23
)

func main() {
	fmt.Printf("O tipo de lucro é %T", lucro)
}
