package main

import (
	"fmt"
)

var x int = 7
var g func() = func() {
	fmt.Println("g desde fuera")
}

func main() {
	f := func() {
		for i := 0; i <= 3; i++ {
			fmt.Println(i)
		}
	}

	f()

	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", g)
	fmt.Println(x)
	g()
}
