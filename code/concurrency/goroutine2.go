package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("Hello")
}

func main() {
	printHello := func() {
		fmt.Println("Hello from goroutine")
	}

	go printHello() // HL
	fmt.Println("main function")
	time.Sleep(1 * time.Second)
}
