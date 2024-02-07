package main

import "fmt"

func main() {
	var ch chan int // HL
	fmt.Println(ch)
	fmt.Println(ch == nil)

	c := make(chan int) // HL
	fmt.Println(c)
}
