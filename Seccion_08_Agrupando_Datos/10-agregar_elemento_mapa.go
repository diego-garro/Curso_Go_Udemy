package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
	}

	fmt.Println(m)
	fmt.Println(m["Batman"])
	fmt.Println(m["Diego"])
	v, ok := m["Diego"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Diego"]; ok {
		fmt.Println("No se imprime nada", v)
	}

	if v, ok := m["Robin"]; ok {
		fmt.Println("Imprimiendo desde el if", v)
	}

	m["Diego"] = 33
	fmt.Println(m["Diego"])

	for llave, valor := range m {
		fmt.Println(llave, valor)
	}
}
