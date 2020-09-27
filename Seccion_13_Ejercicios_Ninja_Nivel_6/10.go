package main

import (
	"fmt"
)

func main() {
	f := woo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func woo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
