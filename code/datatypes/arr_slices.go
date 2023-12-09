package main

import "fmt"

func main() {
	names := []string{"Rajesh", "Ramesh", "Suresh"} // HL
	names = append(names, "Rajesh", "Ramesh", "Suresh")
	fmt.Println("names:", names, "Len:", len(names), "Cap:", cap(names))

	n := names[1:3] // HL
	fmt.Println("n:", n, "Len:", len(n), "Cap:", cap(n))

	n[1] = "RAJESH" // HL
	fmt.Println("names:", names)
	fmt.Println("n:", n)

	// strings are immutable, inplace changes are not allowed
	s := "SWECHA"
	// s[1] = 'W' // HL this line will not compile
	fmt.Println("s:", s, "Len:", len(s))
	fmt.Println("Accessing three characters:", s[0:3])
}
