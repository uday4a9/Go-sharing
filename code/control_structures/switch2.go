package main

import "fmt"

func main() {
	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); { // HL
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}
}
