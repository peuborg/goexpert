package main

import (
	"fmt"
)

func main() {
	//Closures - Funções Anônimas
	total := func() float64 {
		return somar(50, 10, 56, 89, 4, 546, 59, 840, 65, 23, 444) * 2
	}()
	fmt.Println(total)

}

func somar(valores ...float64) float64 {
	soma := 0.0
	for _, v := range valores {
		soma += v
	}
	return soma
}
