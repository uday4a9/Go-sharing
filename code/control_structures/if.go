package main

import "fmt"

func main() { // HL
	// here n is a local variable in main function
	n := 10 // HL

	if n > 0 { // HL
		fmt.Println("n is positive")
	} else if n < 0 { // HL
		fmt.Println("n is negative")
	} else { // HL
		fmt.Println("n is zero")
	}

}
