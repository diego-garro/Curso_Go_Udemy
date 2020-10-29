package main

import "fmt"

func main() {
	// buffered channel (canal con búfer)
	cs := make(chan<- int, 2) //send only chan<-
	cr := make(<-chan int, 2) //receive only <-chan
	c := make(chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("---------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
}
