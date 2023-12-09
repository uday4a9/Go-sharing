package main

import "fmt"

type Employee struct {
	Id   int
	Name string
}

type Employee2 struct {
	Id   int
	Name string
}

func main() {

	// 1. Declare a struct variable
	var emp Employee // HL
	emp.Id = 1
	emp.Name = "John Doe"
	fmt.Println("emp:", emp)

	// 2. Declare and initialize a struct variable
	emp2 := Employee{Id: 2, Name: "Jane Doe"} // HL
	fmt.Println("emp2:", emp2)

	// Comparing the struct variable
	fmt.Println("emp == emp2:", emp == emp2) // HL

	// 3. Declare and initialize a struct2, and compare with struct1
	// emp3 := Employee2{Id: 2, Name: "Jane Doe"} // HL
	// fmt.Println("emp2 == emp3:", emp2 == emp3) // HL
}
