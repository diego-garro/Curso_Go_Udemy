package main

import (
	"fmt"
)

func main() {
	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		} else if len(xi) == 1 {
			return xi[0]
		} else {
			return xi[0] + xi[len(xi)-1]
		}
	}

	x := bar1(g, []int{1, 2, 3, 4, 5})
	fmt.Println(x)
}

func bar1(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}
