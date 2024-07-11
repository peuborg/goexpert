package main

import "fmt"

func main() {
	//Maps
	salarios := map[string]int{"fulano": 10, "beltrano": 20, "ciclano": 30}

	fmt.Println(salarios["beltrano"])

	//delete(salarios, "fulano")

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	sal := make(map[string]int)
	sal["ful"] = 10
}
