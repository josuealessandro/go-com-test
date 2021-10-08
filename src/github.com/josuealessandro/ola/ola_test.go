package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado string, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("Quando o nome é informado", func(t *testing.T) {
		resultado := Ola("Josué")
		esperado := "Olá, Josué"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Quando o nome não é informado", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá, Mundo"

		verificaMensagemCorreta(t, resultado, esperado)

	})

}
