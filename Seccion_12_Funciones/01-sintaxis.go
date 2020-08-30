package main

import (
	"fmt"
)

func main() {
	foo()
	bar("James")
	s1 := woo("MoneyPenny.")
	fmt.Println(s1)
	x, y := saludar("Diego", "Garro")
	fmt.Println(x)
	fmt.Println(y)
}

// func (r receptor) identificador(parámetros) retorno(s) {código}

func foo() {
	fmt.Println("Hola desde foo.")
}

func bar(s string) {
	fmt.Println("Hola,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hola desde woo, ", st)
}

func saludar(n string, a string) (string, bool) {
	p := fmt.Sprint(n, " ", a, ` dice "hola".`)
	q := true
	return p, q
}