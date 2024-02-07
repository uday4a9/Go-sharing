package main

import "fmt"

type Employee struct {
	Name string
}

func main() {
	var s interface{} // HL

	str := "hello"
	s = str // HL
	fmt.Println(s)

	i := 30
	s = i // HL
	fmt.Println(s)

	e := Employee{
		Name: "Bob Bobson",
	}
	s = e // HL
	fmt.Println(s)
}
