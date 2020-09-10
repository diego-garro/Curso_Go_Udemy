package main

import (
	"fmt"
)

func main() {
	foo3()

	func() {
		fmt.Println("La función anónima se ejecutó")
	}()

	func(x int) {
		fmt.Println("El significado de la vida es:", x)
	}(42)
}

func foo3() {
	fmt.Println("foo3 se ejecutó")
}
