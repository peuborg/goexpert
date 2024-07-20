package main

import "fmt"

func main() {
	//em go existe apenas o for
	for i := 0; i < 10; i++ {
		println(i)
	}

	//Slice
	s := []string{"um", "dois", "tres"}

	for i, v := range s {
		fmt.Printf("O valor do índice %d é %s\n", i, v)
	}

	for _, v := range s {
		fmt.Printf("O valor é %s\n", v)
	}

	for i, _ := range s {
		fmt.Printf("O índice é %d\n", i)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	for {
		//loop infinito
		println("loop")
	}
}
