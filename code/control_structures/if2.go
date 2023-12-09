package main

import "fmt"

func main() { // HL
	if n := 10; n%2 == 0 { // HL
		fmt.Println("n is even")
	} else { // HL
		fmt.Println("n is odd")
	}

	// Here n defined within the block of if statement
	// is not accessible outside the block.
	// Accessing it will result in a compile time error.
	// fmt.Println(n) // HL
}
