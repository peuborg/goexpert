package matematica

import "fmt"

//OBS: Structs, Métodos e váriáveis que são exportados devem começar com letra maiúscula
func Soma[T int | int64 | float32 | float64](a T, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}

func (c Carro) Andar() {
	fmt.Printf("Carro %v andou!\n", c.Marca)
}
