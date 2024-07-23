package main

type Numero interface {
	int | int64 | float32 | float64
}

func Soma[T Numero](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	} else {
		return false
	}

}

func main() {
	m := map[string]int{"fulano": 10, "beltrano": 20, "ciclano": 30}
	m2 := map[string]float64{"fulano": 100.1, "beltrano": 200.2, "ciclano": 300.3}
	println(Soma(m))
	println(Soma(m2))
	println(Compara(Soma(m), Soma(m)))
}
