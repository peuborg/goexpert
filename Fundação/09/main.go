package main

import (
	"fmt"
)

func main() {
	//Funções Variáticas
	fmt.Println(somar(50, 10, 56, 89, 4, 546, 59, 840, 65, 23, 444))

}

func somar(valores ...float64) float64 {
	soma := 0.0
	for _, v := range valores {
		soma += v
	}
	return soma
}
