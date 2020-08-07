package main

import (
	"fmt"
)

func main() {
	i := 0
	for i <= 10 {
		fmt.Println("El ciclo sigue...")
		i++
	}

	fmt.Println("Se terminó el ciclo :)")

	for {
		if i > 100 {
			break
		}
		fmt.Println(i)
		i++
	}
}
