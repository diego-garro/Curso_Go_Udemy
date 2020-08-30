package main

import (
	"fmt"
)

func main() {
	s := struct {
		nombre string
		amigos map[string]int
		bebidas []string
	} {
		nombre: "Diego",
		amigos: map[string]int {
			"Carito": 222,
			"Adriana": 444,
			"Condor": 666,
		},
		bebidas: []string {
			"Agua",
			"Chicha",
		},
	}

	fmt.Println(s.nombre)
	fmt.Println(s.amigos)
	fmt.Println(s.bebidas)

	for n, num := range s.amigos {
		fmt.Println(n, num)
	}

	for ind, bebida := range s.bebidas {
		fmt.Println(ind, bebida)
	}
}