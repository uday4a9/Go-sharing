package main

import "fmt"

func main() {
	ch := make(chan int) // HL

	go func() {
		// Write to a channel
		ch <- 12 // HL
	}()

	// Read from a channel
	fmt.Println("Received1:", <-ch) // HL
	// Read More from the same channel
	fmt.Println("Received2:", <-ch) // HL
}
