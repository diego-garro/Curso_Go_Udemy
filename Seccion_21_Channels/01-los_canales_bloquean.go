package main

import "fmt"

func main() {
	// unbuffered channel (canal sin búfer)
	ca := make(chan int)

	ca <- 42
	fmt.Println(<-ca)
}
