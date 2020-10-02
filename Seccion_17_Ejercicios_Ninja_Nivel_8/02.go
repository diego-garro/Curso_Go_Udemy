package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre string `json:"Nombre"`
	Edad   int    `json:"Edad"`
}

func main() {
	s := `[{"Nombre":"Diego","Edad":34},{"Nombre":"Melina","Edad":33},{"Nombre":"Samantha","Edad":5}]`
	fmt.Println(s)

	var personas []persona
	err := json.Unmarshal([]byte(s), &personas)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(personas)

	for i, p := range personas {
		fmt.Println("\tPERSONA #", i+1)
		fmt.Println("\t\t", p.Nombre)
		fmt.Println("\t\t", p.Edad)
	}
}
