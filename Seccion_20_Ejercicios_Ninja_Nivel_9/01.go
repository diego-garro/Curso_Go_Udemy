package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("Número de CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Número de Goroutines: %v\n", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hola desde la primera Go rutina")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hola desde la segunda Go rutina")
		wg.Done()
	}()

	fmt.Printf("Número de CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Número de Goroutines: %v\n", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("A punto de finalizar main.")
	fmt.Printf("Número de CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Número de Goroutines: %v\n", runtime.NumGoroutine())
}
