package main

import (
	"fmt"
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

}
