package main

import (
	"fmt"
)

type persona1 struct {
	nombre   string
	apellido string
}

type agenteSecreto1 struct {
	persona1
	lpm bool
}

func (a agenteSecreto1) presentar() {
	fmt.Println("Hola, soy", a.nombre, a.apellido)
}

func (p persona1) presentar() {
	fmt.Println("Hola, soy", p.nombre, p.apellido, "la persona se presenta")
}

type humano interface {
	presentar()
}

func bar2(h humano) {
	switch h.(type) {
	case persona1:
		fmt.Println("Fui pasado a la función bar2 (persona1)", h.(persona1).nombre)
	case agenteSecreto1:
		fmt.Println("Fui pasado a la función bar2 (agenteSecreto1)", h.(agenteSecreto1).nombre)
	}
	fmt.Println("Fui pasado a la funcion bar2", h)
}

type manzana int

func main() {
	as3 := agenteSecreto1{
		persona1: persona1{
			nombre:   "Diego",
			apellido: "Garro",
		},
		lpm: true,
	}

	as4 := agenteSecreto1{
		persona1: persona1{
			nombre:   "Melina",
			apellido: "Madrigal",
		},
		lpm: false,
	}

	p := persona1{
		nombre:   "Samantha",
		apellido: "Garro",
	}

	fmt.Println(as3)
	fmt.Println(as4)
	as3.presentar()
	as4.presentar()

	bar2(as3)
	bar2(as4)
	bar2(p)

	// Conversión
	var x manzana = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
