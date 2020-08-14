package main

import (
	"fmt"
)

func main() {
	x := "Robin"
	if x == "Batman" {
		fmt.Println(x)
	} else if x == "Robin" {
		fmt.Println("Santas cachuchas Batman")
	} else {
		fmt.Println("Ninguno")
	}
}
