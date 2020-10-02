package main

import (
	"encoding/json"
	"fmt"
)

type usuario struct {
	Nombre string
	Edad   int
}

func main() {
	u1 := usuario{
		Nombre: "Diego",
		Edad:   34,
	}
	u2 := usuario{
		Nombre: "Melina",
		Edad:   33,
	}
	u3 := usuario{
		Nombre: "Samantha",
		Edad:   5,
	}

	usuarios := []usuario{u1, u2, u3}
	fmt.Println(usuarios)

	bs, err := json.Marshal(usuarios)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
