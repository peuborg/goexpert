package main

import "fmt"

var (
	conectou bool    = true
	tempo    int     = 10
	resposta string  = "sucesso"
	lucro    float64 = 500.23
)

func main() {
	//Array
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(len(meuArray))

	for i, v := range meuArray {
		fmt.Printf("O valor do índice %d é %d\n", i, v)
	}
}
