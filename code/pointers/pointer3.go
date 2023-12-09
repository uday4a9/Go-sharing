package main

import "fmt"

func main() {
	type Employee struct { // HL
		Name string
		Age  int
	}

	// This points to an address which points to zero vaule of the type
	eptr := new(Employee) // HL
	fmt.Println("eptr:", eptr)

	var e Employee = Employee{"John", 30}
	var eptr1 *Employee = &e // HL
	fmt.Println("Name:", (*eptr1).Name, "Age:", (*eptr1).Age)
	fmt.Println("Name:", eptr1.Name, "Age:", eptr1.Age)
}
