package main

import (
	"fmt"
)

type persona struct {
	nombre   string
	apellido string
}

type agenteSecreto struct {
	persona
	lpm bool
}

func (a agenteSecreto) presentarse() {
	fmt.Println("Hola, soy", a.nombre, a.apellido)
}

func main() {
	as1 := agenteSecreto{
		persona: persona{
			nombre:   "Diego",
			apellido: "Garro",
		},
		lpm: true,
	}

	as2 := agenteSecreto{
		persona: persona{
			nombre:   "Melina",
			apellido: "Madrigal",
		},
		lpm: false,
	}

	fmt.Println(as1)
	fmt.Println(as2)
	as1.presentarse()
	as2.presentarse()
}

// func (r receptor) identificador(parámetros) (retorno) {código}
