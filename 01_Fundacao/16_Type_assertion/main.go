package main

import "fmt"

func main() {
	var minhaVar interface{} = "Peuborg"
	println(minhaVar.(string))
	valorConvertido, resultadoConversao := minhaVar.(int)
	fmt.Printf("O valorConvertido é %v e o resultadoConversao é %v", valorConvertido, resultadoConversao)
}
