package main

import (
	"encoding/json"
	"fmt"
)

type persona1 struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
}

func main() {
	s := `[{"Nombre":"James","Apellido":"Bond","Edad":32},{"Nombre":"Miss","Apellido":"Moneypenny","Edad":27}]`
	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//personas1 = []persona1{}
	var personas1 []persona1

	err := json.Unmarshal(bs, &personas1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Toda la data", personas1)
	for i, v := range personas1 {
		fmt.Println("\nPERSONA NÚMERO", i+1)
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
	}
}
