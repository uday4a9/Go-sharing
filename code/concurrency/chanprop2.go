package main

import "fmt"

func main() {
	ch := make(chan int) // HL
	fmt.Println(ch)

	// Read from a channel
	fmt.Println(<-ch) // HL
}
