package main

import (
	"fmt"
)

func main() {
	var arr [3]int // HL
	fmt.Println("arr:", arr)

	arr1 := [3]int{1, 2, 3} // HL
	fmt.Println("arr1:", arr1)

	// Sparse array intialization
	arr2 := [5]string{"Hello", 3: "swecha"} // HL
	fmt.Println("arr2:", arr2)
	fmt.Println("Length:", len(arr2))
}
