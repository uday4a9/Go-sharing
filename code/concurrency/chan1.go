package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)     // HL
	go func(ch chan<- int) { // HL
		fmt.Println("Sending:", 1)
		time.Sleep(2 * time.Second)
		ch <- 1
	}(ch) // HL
	fmt.Println("Received:", <-ch) // HL
}
