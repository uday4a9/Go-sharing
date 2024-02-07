package main

import "fmt"

func main() {
	ch := make(chan int) // HL

	go func() {
		// Write to a channel
		ch <- 12 // HL
	}()

	// Read from a channel
	fmt.Println("Received", <-ch) // HL
}
