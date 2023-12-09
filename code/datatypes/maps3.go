package main

import "fmt"

func main() {
	// 1. make a map
	capitalMap := make(map[string]string) // HL
	capitalMap["AP"] = "Amaravati"
	capitalMap["TS"] = "Hyderabad"
	capitalMap["KA"] = "Bengaluru"

	fmt.Println("nmap:", capitalMap, "Len:", len(capitalMap))

	// Accessing Non-existent key, Hw do we know if the key exists?
	capital, ok := capitalMap["AP"] // HL
	fmt.Println("capital:", capital, "ok:", ok)
	// capital, ok = capitalMap["MH"] // HL
	// fmt.Println("capital:", capital, "ok:", ok)
}
