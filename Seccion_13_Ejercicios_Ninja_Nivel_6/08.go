package main

import (
	"fmt"
)

func main() {
	f := foo1()
	fmt.Println(f())
}

func foo1() func() int {
	return func() int {
		return 42
	}
}
