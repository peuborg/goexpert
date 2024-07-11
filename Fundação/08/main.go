package main

import (
	"errors"
	"fmt"
)

func main() {
	//Funções
	resultado, err := dividir(50, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}

func dividir(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0, errors.New("Divisão por zero")
	} else {
		return a / b, nil
	}
}
