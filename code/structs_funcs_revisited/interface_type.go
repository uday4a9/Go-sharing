package main

import "fmt"

type Printer interface {
	Print()
}

type Employee struct {
	Name string
}

func (e Employee) Print() {
	fmt.Println(e.Name)
}

type Person struct {
	Name string
}

func (p Person) Print() {
	fmt.Println(p.Name)
}

func Print(p Printer) {
	p.Print()
}

func main() {

	e := Employee{
		Name: "Bob Bobson",
	}
	Print(e)

	p := Person{
		Name: "Alice Alisson",
	}

	Print(p)
}
