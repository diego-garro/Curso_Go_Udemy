package main

import (
	"fmt"
)

func main() {
	s1 := foo4()
	fmt.Println(s1)

	z := bar3()
	fmt.Printf("%T\n", z)

	i := z()
	fmt.Println(i)
}

func foo4() string {
	s := "Hola mundo"
	return s
}

func bar3() func() int {
	return func() int {
		return 1492
	}
}
