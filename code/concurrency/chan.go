package main

import "fmt"

func main() {
	var ch chan int
	fmt.Println(ch == nil) // HL

	// ch := make(chan int) // HL
	// go func() {
	// 	ch <- 1 // HL
	// }()
	// fmt.Println(<-ch) // HL
}
