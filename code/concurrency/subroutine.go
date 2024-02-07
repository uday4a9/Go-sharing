package main

import (
	"fmt"
)

func printHello() { // HL
	fmt.Println("Hello")
}

func main() {
	printHello() // HL
	fmt.Println("main function")
}
