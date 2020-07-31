package main

import (
	"fmt"
)

type mi_tipo int

var x mi_tipo
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("El tipo de la variable x es: %T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("El tipo de la variable y es: %T\n", y)
}
