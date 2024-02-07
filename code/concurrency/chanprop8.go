package main

import "fmt"

func gen(ch chan<- int) {
	defer close(ch) // HL
	for i := 0; i < 5; i++ {
		ch <- i * 2 // HL
	}
}

func main() {
	ch := make(chan int)
	go gen(ch)
	for i := range ch { // HL
		fmt.Println(i)
	}
}
