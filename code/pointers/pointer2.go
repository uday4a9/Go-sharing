package main

import "fmt"

func main() {
	// This points to an address which points to zero vaule of the type
	nptr := new(int) // HL
	fmt.Println("nptr:", nptr)
	fmt.Println("*nptr:", *nptr)
	*nptr = 10
	fmt.Println("*nptr:", *nptr)
}
