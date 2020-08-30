package main

import (
	"fmt"
)

type persona struct {
	nombre string
	apellido string
	saboresFav []string
}

func main() {
	p1 := persona {
		nombre: "Diego",
		apellido: "Garro",
		saboresFav: []string {
			"chocolate",
			"combinado",
			"galleta",
		},
	}

	p2 := persona {
		nombre: "Melina",
		apellido: "Madrigal",
		saboresFav: []string {
			"Fresa",
			"Vainilla",
			"Chocochips",
		},
	}
	
	fmt.Println(p1.nombre)
	fmt.Println(p1.apellido)
	for i, v := range p1.saboresFav {
		fmt.Println("\t", i, v)
	}

	fmt.Println(p2.nombre)
	fmt.Println(p2.apellido)
	for i, v := range p2.saboresFav {
		fmt.Println("\t", i, v)
	}
}