package lib

import (
	"testing"
)

func TestSubstrair(t *testing.T) {
	resultado := Subtrair(2, 1)
	if resultado != 1 {
		t.Errorf("A funcao Substrair (pacote lib) com parametro 2 1 retornou = %d", resultado)
	}
}
