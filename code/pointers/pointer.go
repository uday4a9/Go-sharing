package main

import "fmt"

func main() {
	var intPtr *int // HL
	fmt.Println("intPtr:", intPtr)
	// dereference a nil pointer, leads to panic
	// fmt.Println("intPtr:", *intPtr) // HL
	n := 10
	intPtr = &n // HL
	fmt.Println("intPtr:", *intPtr)

	*intPtr = 20 // HL
	fmt.Println("intPtr:", *intPtr, "n:", n)
}
