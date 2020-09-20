package main

import (
	"fmt"
	"math"
)

type circulo struct {
	radio float64
}

type cuadrado struct {
	longitud float64
}

func (c circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c cuadrado) area() float64 {
	return c.longitud * c.longitud
}

type forma interface {
	area() float64
}

func info(f forma) {
	fmt.Println(f.area())
}

func main() {
	circ := circulo{
		radio: 12.345,
	}

	cuad := cuadrado{
		longitud: 15.1,
	}

	info(circ)
	info(cuad)
}
