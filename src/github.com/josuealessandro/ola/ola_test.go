package main

import "testing"

func TestOla(t *testing.T) {
	t.Run("Quando o nome é informado", func(t *testing.T) {
		resultado := Ola("Josué")
		esperado := "Olá, Josué"

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("Quando o nome não é informado", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá, Mundo"

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

}
