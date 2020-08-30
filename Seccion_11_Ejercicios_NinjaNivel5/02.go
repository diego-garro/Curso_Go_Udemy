package main

import (
	"fmt"
)

type persona1 struct {
	nombre string
	apellido string
	saboresFav []string
}

func main() {
	p11 := persona1 {
		nombre: "Diego",
		apellido: "Garro",
		saboresFav: []string {
			"chocolate",
			"combinado",
			"galleta",
		},
	}

	p22 := persona1 {
		nombre: "Melina",
		apellido: "Madrigal",
		saboresFav: []string {
			"Fresa",
			"Vainilla",
			"Chocochips",
		},
	}

	m := map[string]persona1 {
		p11.apellido: p11,
		p22.apellido: p22,
	}
	
	for _, val := range m {
		fmt.Println(val.nombre)
		fmt.Println(val.apellido)
		for index, value := range val.saboresFav {
			fmt.Println("\t", index, value)
		}
		fmt.Println("---------------------")
	}
}