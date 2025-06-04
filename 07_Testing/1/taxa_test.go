package taxa

import "testing"

func TestCalcularTaxa(t *testing.T) {
	montante := 500.0
	expectativa := 5.0

	resultado := CalcularTaxa(montante)

	if resultado != expectativa {
		t.Errorf("Expectativa %f mas obtivemos %f", expectativa, resultado)
	}
}

func TestCalcularTaxaBatch(t *testing.T) {
	type calcTax struct {
		montante, expectativa float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}

	for _, item := range table {
		resultado := CalcularTaxa(item.montante)

		if resultado != item.expectativa {
			t.Errorf("Expectativa %f mas obtivemos %f", item.expectativa, resultado)
		}
	}

}

func BenchmarkCalcularTaxa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalcularTaxa(500.0)
	}
}

func FuzzCalcularTaxa(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}

	for _, montante := range seed {
		f.Add(montante)
	}

	f.Fuzz(func(t *testing.T, montante float64) {
		resultado := CalcularTaxa(montante)

		if montante <= 0 && resultado != 0 {
			t.Errorf("Recebido %f mas esparava 0", resultado)
		}
	})
}
