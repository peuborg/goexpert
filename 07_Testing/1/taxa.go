package taxa

func CalcularTaxa(montante float64) float64 {
	if montante == 0.0 {
		return 0.0
	}
	if montante >= 1000 {
		return 10.0
	}
	return 5.0
}
