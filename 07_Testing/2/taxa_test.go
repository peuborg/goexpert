package taxa

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCalcularTaxa(t *testing.T) {

	taxa, err := CalcularTaxa(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, taxa)

	taxa, err = CalcularTaxa(0.0)
	assert.Error(t, err, "Montante deve ser maior que 0")
	assert.Equal(t, 0.0, taxa)
}

func TestCalcularTaxaESalvar(t *testing.T) {
	repositorio := &TaxaRepositorioMock{}
	repositorio.On("SalvarTaxa", 10.0).Return(nil)
	repositorio.On("SalvarTaxa", 0.0).Return(errors.New("Erro ao salvar taxa"))
	repositorio.On("SalvarTaxa", mock.Anything).Return(errors.New("Erro ao salvar taxa"))

	err := CalcularTaxaESalvar(1000.0, repositorio)
	assert.Nil(t, err)

	err = CalcularTaxaESalvar(0.0, repositorio)
	assert.Error(t, err, "Erro ao salvar taxa")

	repositorio.AssertExpectations(t)
}
