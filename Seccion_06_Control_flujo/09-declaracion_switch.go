package main

import (
	"fmt"
)

func main() {
	switch {
	case 2 == 4, 4 == 8, 9 == 9:
		fmt.Println("No debería imprimir esto, ahora sí :)")
	case 3 == 3:
		fmt.Println("Esto debería imprimirlo")
		fallthrough
	case 4 == 5:
		fmt.Println("Esto tampoco debería imprimirlo")
	default:
		fmt.Println("Imprimiendo desde default")
	}

	switch "Manzana" {
	case "Limon", "Pera":
		fmt.Println("Frutas verdes")
	case "Manzana", "Fresa":
		fmt.Println("Frutas rojas")
		fallthrough
	default:
		fmt.Println("Caso por defecto")
	}
}
