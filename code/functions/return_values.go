package main

import (
	"errors"
	"fmt"
)

func divAndRemainder(numerator int, denominator int) (int, int, error) { // HL
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	result, remainder, err := divAndRemainder(5, 0) // HL
	if err != nil {
		fmt.Println("err1:", err)
	} else {
		fmt.Println(result, remainder)
	}

	// We can ignore the remainder by using the blank identifier
	result, _, err = divAndRemainder(5, 2) // HL
	if err != nil {
		fmt.Println("err2:", err)
	} else {
		fmt.Println(result)
	}
}
