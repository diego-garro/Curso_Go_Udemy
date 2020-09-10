package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum := suma(ii...)
	fmt.Println(sum)

	sum1 := sumaPares(suma, ii...)
	fmt.Println("El total de los pares es:", sum1)

}

func suma(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}

	return total
}

func sumaPares(f func(xi ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}
	}

	return f(y...)
}
