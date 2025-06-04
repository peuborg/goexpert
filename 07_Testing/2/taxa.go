package taxa

import "errors"

func CalcularTaxa(montante float64) (float64, error) {
	if montante <= 0.0 {
		return 0.0, errors.New("Montante deve ser maior que 0")
	}
	if montante == 0.0 {
		return 0.0, nil
	}
	if montante >= 1000 {
		return 10.0, nil
	}
	return 5.0, nil
}

func CalcularTaxa2(montante float64) float64 {
	if montante <= 0.0 {
		return 0.0
	}
	if montante >= 1000 && montante < 20000 {
		return 10.0
	}
	if montante >= 20000 {
		return 20.0
	}
	return 5.0
}

type Repositorio interface {
	SalvarTaxa(taxa float64) error
}

func CalcularTaxaESalvar(montante float64, repositorio Repositorio) error {
	taxa := CalcularTaxa2(montante)
	return repositorio.SalvarTaxa(taxa)
}
