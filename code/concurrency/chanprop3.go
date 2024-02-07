package main

import "fmt"

func main() {
	ch := make(chan int) // HL
	fmt.Println(ch)

	// Write to a channel
	ch <- 12 // HL
}
