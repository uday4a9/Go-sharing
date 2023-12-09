package main

import "fmt"

type Employee struct {
	Id   int
	Name string
}

func main() {
	// 1. Declare a struct variable
	var employees []Employee // HL
	employees = append(employees, Employee{Id: 1, Name: "John Doe"})
	employees = append(employees, Employee{Id: 2, Name: "Jane Doe"})
	employees = append(employees, Employee{Id: 3, Name: "Jack Doe"})

	fmt.Println("employees:", employees, "Len:", len(employees), "Cap:", cap(employees))
}
