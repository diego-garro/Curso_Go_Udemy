package main

import "fmt"

func main() {
	dg := []string{"Diego", "Garro", "Fútbol", "Beisbol", "Correr"}
	mm := []string{"Melina", "Madrigal", "Nadar", "Montañismo", "Basket"}

	fmt.Println(dg)
	fmt.Println(mm)

	vp := [][]string{dg, mm}
	fmt.Println(vp)
}
