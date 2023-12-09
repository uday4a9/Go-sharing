package main

import "fmt"

func main() { // HL
	names := []string{"John", "Paul", "George", "Ringo"}
	for i, name := range names { // HL
		fmt.Println(i, name)
	}

	nmap := map[string]int{"John": 1, "Paul": 2, "George": 3, "Ringo": 4}
	for name, number := range nmap { // HL
		fmt.Println(name, number)
	}
}
