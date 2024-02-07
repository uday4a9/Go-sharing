package main

import "fmt"

func main() {
	// created a channel and closed it
	ch := make(chan int)
	go func() {
		defer close(ch) // HL
		ch <- 12        // HL
		ch <- 13        // HL
	}()

	fmt.Println("Received1:", <-ch) // HL
	fmt.Println("Received2:", <-ch) // HL
}
