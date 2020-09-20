package main

import (
	"fmt"
)

func main() {
	defer boo()
	fmt.Println("Hola desde main")
}

func boo() {
	defer func() {
		fmt.Println("boo diferida corrió")
	}()
	fmt.Println("boo corrió")
}
