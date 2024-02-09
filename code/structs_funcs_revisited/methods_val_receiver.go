package main

import "fmt"

type Employee struct {
	Name string
	Age  int
}

// Display() is a method : Value receiver
func (e Employee) Display() { // HL
	fmt.Println(e.Name, "is", e.Age, "years old.")
}

// SetAge() is a method : Value receiver
func (e Employee) SetAge(age int) { // HL
	e.Age = age
}

func main() {
	e := Employee{"John Doe", 25}
	e.Display()
	// e.SetAge(30)
	// e.Display()
}
