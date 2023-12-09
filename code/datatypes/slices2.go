package main

import "fmt"

func main() {
	capitals := make([]string, 2, 5) // HL
	fmt.Println("len:", len(capitals), "cap:", cap(capitals))

	capitals = append(capitals, "Hyderabad", "Bengaluru", "Chennai") // HL
	fmt.Println("len:", len(capitals), "cap:", cap(capitals))

	// capitals = append(capitals, "Mumbai")               // HL
	// capitals = append(capitals, "Delhi", "Bhubaneswar") // HL
	// fmt.Println("len:", len(capitals), "cap:", cap(capitals))

	// zero length slices
	// names := make([]string, 0, 5) // HL
	// names = append(names, "Rajesh", "Ramesh", "Suresh")
	// fmt.Println("len:", len(names), "cap:", cap(names))
}
