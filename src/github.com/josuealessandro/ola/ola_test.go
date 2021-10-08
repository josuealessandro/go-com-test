package main

import "testing"

func TestOla(t *testing.T){
	resultado := Ola("Josué")
	esperado := "Olá, Josué"

	if(resultado != esperado){
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
