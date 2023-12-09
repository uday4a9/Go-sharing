package main

import "fmt"

func main() {
	capitals := [5]string{"Hyderabad", "Bengaluru", "Chennai"} // HL
	fmt.Println("names:", capitals)

	// Accessing array elements 0 to len(arr)-1
	fmt.Println("names[0]:", capitals[0])

	// Accesing beyond range elements leads to panic
	// fmt.Println("names[5]:", capitals[5])

	// southCapitals := [3]string{"Hyderabad", "Bengaluru", "Chennai"} // HL
	// Operators on arrays: Equality and Non-equality

	// fmt.Println("capitals == southCapitals:", capitals == southCapitals) // HL
}
