package main

import "fmt"

func main() {

	func() { // HL
		fmt.Println("Hello, anonymous function!")
	}() // HL

	fmt.Println("Hello, playground")
}
