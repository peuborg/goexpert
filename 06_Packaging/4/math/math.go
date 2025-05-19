package math

//Estruturas iniciadas com letra maiúscula ficam com escopo público
type Math struct {
	//Variáveis iniciadas com letra maiúscula ficam com escopo público
	A int
	B int
}

//Funções iniciadas com letra maiúscula ficam com escopo público
func (m Math) Add() int {
	return m.A + m.B
}

//Estruturas iniciadas com letra minúscula ficam com escopo interno (private)
type matematica struct {
	//Variáveis iniciadas com letra minúscula ficam com escopo interno (private)
	a int
	b int
}

//Funções iniciadas com letra minúscula ficam com escopo interno (private)
func NewMath(a, b int) matematica {
	return matematica{a: a, b: b}
}
