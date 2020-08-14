package main

import (
	"fmt"
)

func main() {
	// Literal compuesto (Composite literal)
	x := []int{2, 3, 4, 7, 42}
	fmt.Println(x)
	fmt.Println(cap(x))

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println(x[1:])
}
