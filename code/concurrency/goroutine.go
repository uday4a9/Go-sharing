package main

import (
	"fmt"
)

func printHello() {
	fmt.Println("Hello")
}

func main() {
	go printHello() // HL
	fmt.Println("main function")
}
