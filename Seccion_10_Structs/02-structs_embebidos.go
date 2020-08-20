package main

import (
	"fmt"
)

type person struct {
	nombre   string
	apellido string
	edad     int
}

type agenteSecreto struct {
	person
	lpm bool
}

func main() {
	as1 := agenteSecreto{
		person: person{
			nombre:   "James",
			apellido: "Bond",
			edad:     40,
		},
		lpm: true,
	}

	p2 := person{
		nombre:   "Melina",
		apellido: "Madrigal",
		edad:     33,
	}

	fmt.Println(as1)
	fmt.Println(p2)

	fmt.Println(as1.nombre, as1.apellido, as1.edad)
	fmt.Println(p2.nombre, p2.apellido, p2.edad)
}
