package main

import (
	"fmt"
)

func main() {
	defer woo1()
	bar1()
}

func woo1() {
	fmt.Println("Woo")
}

func bar1() {
	fmt.Println("Bar")
}
