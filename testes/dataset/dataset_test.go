package dataset

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - indice: esperado(%d) <> encontrado(%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"GoLang é Show", "GoLang", 0},
		{"", "a", 0},
		{"Ola", "ola", -1},
		{"Carlos", "r", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}

}
