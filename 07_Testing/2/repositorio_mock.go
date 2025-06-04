package taxa

import (
	"github.com/stretchr/testify/mock"
)

type TaxaRepositorioMock struct {
	mock.Mock
}

func (m *TaxaRepositorioMock) SalvarTaxa(taxa float64) error {
	args := m.Called(taxa)
	return args.Error(0)
}
