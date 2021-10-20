package lib

import (
	"testing"
)

func TestAdicionar(t *testing.T) {
	resultado := Adicionar(1, 2)
	if resultado != 3 {
		t.Errorf("Adicionar(1, 2) retornou %d", resultado)
	}
}
