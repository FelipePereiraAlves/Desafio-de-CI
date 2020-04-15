package desafio_ci_golang_soma

import (
	"testing"
)

func TestSoma(t *testing.T) {
	actualResult := soma(5, 5)
	var expectedResult = 10

	if actualResult != expectedResult {
		t.Errorf("Soma incorreta, obtido: %d, esperado: %d.", actualResult, expectedResult)
	}
	
}