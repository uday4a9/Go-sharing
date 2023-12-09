package main

import "fmt"

func main() {
	i := 1
	for { // HL

		if i > 10 {
			break
		}
		fmt.Println("Hello")
		i++
	}
}
