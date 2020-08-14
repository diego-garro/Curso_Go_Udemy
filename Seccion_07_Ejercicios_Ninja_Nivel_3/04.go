package main

import (
	"fmt"
)

func main() {
	cu := 1986
	for {
		if cu > 2020 {
			break
		}
		fmt.Println(cu)
		cu++
	}
}
