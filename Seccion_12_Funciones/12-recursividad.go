package main

import (
	"fmt"
)

func main() {
	fmt.Println(4 * 3 * 2 * 1)

	n1 := factorial(4)
	fmt.Println(n1)

	n2 := cicloFactorial(4)
	fmt.Println(n2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 4 * factorial(3)...

// Con ciclos
func cicloFactorial(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
