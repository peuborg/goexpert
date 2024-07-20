package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "alô mundo"

	exibirTipo(x)
	exibirTipo(y)
}

func exibirTipo(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é [%v]\n", t, t)
}
