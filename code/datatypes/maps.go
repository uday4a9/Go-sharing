package main

import "fmt"

func main() {
	// 1. Just declare a map
	var capitalMap map[string]string // HL
	fmt.Println("nmap:", capitalMap)

	// Accessing and updating map elements
	fmt.Println(`capitalMap["AP"]:`, capitalMap["AP"])
	// capitalMap["AP"] = "Amaravati" // HL

}
