package main

import (
	"fmt"
)

func main() {
	deporteFav := "fútbol"
	switch deporteFav {
	case "beisbol":
		fmt.Println("ve al estadio")
	case "natación":
		fmt.Println("ve a la playa")
	case "fútbol":
		fmt.Println("ve a la cancha")
	}
}
