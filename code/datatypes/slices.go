package main

import "fmt"

func main() {
	var names []string // HL
	fmt.Println("names:", names)

	capitals := []string{"Hyderabad", "Bengaluru", "Chennai"} // HL
	fmt.Println("names:", capitals)
	fmt.Println("len:", len(capitals), "cap:", cap(capitals))

	// Accessing slice elements similar to arrays
	// Accesing beyond range elements leads to panic

	// southCapitals := [3]string{"Hyderabad", "Bengaluru", "Chennai"} // HL
	// Operators on arrays: Equality and Non-equality

	// fmt.Println("capitals == southCapitals:", capitals == southCapitals) // HL
}
