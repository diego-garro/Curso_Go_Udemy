package main

import (
	"fmt"
)

// type persona1 struct {
// 	first string
// 	last  string
// 	age   int
// }

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Diego",
		last:  "Garro",
		age:   34,
	}

	fmt.Println(p1)
}
