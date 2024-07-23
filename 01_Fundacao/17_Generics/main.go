package main

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{"fulano": 10, "beltrano": 20, "ciclano": 30}
	m2 := map[string]float64{"fulano": 100.1, "beltrano": 200.2, "ciclano": 300.3}
	println(SomaInteiro(m))
	println(SomaFloat(m2))
}
