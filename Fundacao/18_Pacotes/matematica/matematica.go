package matematica

func Soma[T int | int64 | float32 | float64](a T, b T) T {
	return a + b
}
