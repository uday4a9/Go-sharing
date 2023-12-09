package main

import "fmt"

func main() { // HL
	m := map[string]int{
		"one": 1,
	}

	if v, ok := m["one"]; ok { // HL
		fmt.Println("Retrieved value from map for 'one' is:", v)
	}

	if _, ok := m["two"]; !ok { // HL
		fmt.Println("two not found")
	}

}
