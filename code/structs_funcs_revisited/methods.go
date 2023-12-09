package main

import "fmt"

type Employee struct {
	Name string
	Age  int
}

// String() is a method : Value receiver
func (e Employee) String() string {
	return fmt.Sprintf("%s (%d years)", e.Name, e.Age)
}

// SetAge() is a method : Pointer receiver
func (e Employee) SetAge(age int) {
	e.Age = age
}

func main() {
	fmt.Println("Hello, World!")

	e := Employee{"John Doe", 25}
	fmt.Println(e.String())
	// e.SetAge(30)
	// fmt.Println(e.String())

}
