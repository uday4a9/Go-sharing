package main

import "fmt"

func main() {
	foo := func() { // HL
		fmt.Println("Hello, anonymous function (AKA foo)!")
	}

	foo() // HL
	// defer foo() // HL
	fmt.Println("Hello, playground")
}
