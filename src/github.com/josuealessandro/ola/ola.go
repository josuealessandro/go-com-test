package main

import "fmt"

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return PrefixodeSaudacao(idioma) + nome
}

func PrefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case "frances":
		prefixo = prefixoOlaFrances
	case "espanhol":
		prefixo = prefixoOlaEspanhol
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
