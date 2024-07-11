package main

import "fmt"

func main() {
	//Slice
	s := []int{10, 20, 30, 40, 50}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 60) //dobra o tamanho do slide e adiciona no final

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
