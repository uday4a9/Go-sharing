package main

import "fmt"

type Employee struct {
	Name string
	Id   int
}

func DisplayEmployee(e Employee) {
	fmt.Println("Employee Name:", e.Name, "and Id:", e.Id)
}

func ChangeEmployeeId(e Employee, newId int) { // HL
	e.Id = newId
}

func main() {
	e := Employee{"John Doe", 25}
	DisplayEmployee(e)
	// ChangeEmployeeId(e, 30) // HL
	// DisplayEmployee(e)
}
