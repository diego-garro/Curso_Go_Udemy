package main

import (
	"fmt"
)

func main() {
	// for init, condition, post
	for i := 0; i <= 100; i++ {
		fmt.Println("Iteración:", i)
		for j := 0; j <= 10; j++ {
			fmt.Println("\t\titeración: ", j)
		}
	}
}
