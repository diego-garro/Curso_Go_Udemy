package main

import (
	"fmt"
)

func main() {
	x := map[string][]string{
		`eduar_tua`:    []string{`computadoras`, `montaña`, `playa`},
		`carlos_ramon`: []string{`leer`, `comprar`, `música`},
		`juan_bimba`:   []string{`helado`, `pintar`, `bailar`},
	}

	x[`luis_lira`] = []string{`trabajar`, `playa`, `cerveza`}

	delete(x, `eduar_tua`)

	for llave, valor := range x {
		fmt.Println("Registro:", llave)
		for i, val := range valor {
			fmt.Println("\t", i, val)
		}
	}
}
