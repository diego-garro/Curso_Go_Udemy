package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := woo(ii...)
	fmt.Println(n)

	m := choo(ii)
	fmt.Println(m)
}

func woo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func choo(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
