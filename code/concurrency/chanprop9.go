package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch) // HL

	for i := 1; i <= 4; i++ { // HL
		fmt.Printf("iter %d: %d\n", i, <-ch) // HL
	}
}
