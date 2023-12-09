package main

import "fmt"

func main() {
	// 1. Initialize a map
	capitalMap := map[string]string{
		"AP": "Amaravati",
		"TS": "Hyderabad",
	} // HL
	fmt.Println("nmap:", capitalMap)

	// Accessing and updating map elements
	fmt.Println(`capitalMap["AP"]:`, capitalMap["AP"])
	capitalMap["KA"] = "Bengaluru" // HL
	fmt.Println("nmap:", capitalMap, "Len:", len(capitalMap))

	// Accessing a non-existent key
	fmt.Println(`capitalMap["MH"]:`, capitalMap["MH"])
}
