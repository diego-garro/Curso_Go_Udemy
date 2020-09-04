package main

import (
	"fmt"
)

func main() {
	x := foo1(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("El valor total almacenado en la variable x es:", x)
}

// func (r receptor) identificador(parámetros) (retorno) {código}

func foo1(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	suma := 0
	for i, v := range x {
		suma += v
		fmt.Println("El valor en el índice", i, "suma", v, "al total, quedando", suma)
	}
	fmt.Println("\nEl total es", suma)

	return suma
}
