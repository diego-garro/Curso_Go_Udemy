package main

import (
	"fmt"
)

func main() {
	s1 := "Hola mundo"
	s2 := `Esta es una línea
	partida`
	fmt.Println(s1)
	fmt.Println(s2)

	for i, v := range s1 {
		fmt.Printf("En el índice %d, el valor es: %v\n", i, v)
		fmt.Printf("En el índice %d, el valor es: %#x\n", i, v)
		fmt.Printf("En el índice %d, el valor es: %q\n", i, v)
		fmt.Println("#######################################")
	}
}
