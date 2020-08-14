package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 200)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 201; i < 350; i++ {
		x = append(x, i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
