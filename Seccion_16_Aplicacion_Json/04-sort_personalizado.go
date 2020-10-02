package main

import (
	"fmt"
	"sort"
)

type Persona struct {
	Nombre string
	Edad   int
}

type PorEdad []Persona

func (a PorEdad) Len() int           { return len(a) }
func (a PorEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

type PorNombre []Persona

func (a PorNombre) Len() int           { return len(a) }
func (a PorNombre) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorNombre) Less(i, j int) bool { return a[i].Nombre < a[j].Nombre }

func (p Persona) String() string {
	return fmt.Sprintf("%s: %d", p.Nombre, p.Edad)
}

func main() {
	p1 := Persona{"Eduar", 32}
	p2 := Persona{"Diego", 34}
	p3 := Persona{"Melina", 33}
	p4 := Persona{"Samantha", 5}

	personas2 := []Persona{p1, p2, p3, p4}

	sort.Sort(PorEdad(personas2))
	fmt.Println(personas2)

	sort.Sort(PorNombre(personas2))
	fmt.Println(personas2)
}
