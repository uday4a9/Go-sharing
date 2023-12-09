package main

import "fmt"

type Employee struct {
	Name string
}

func main() {
	var s interface{}

	str := "hello"
	s = str
	fmt.Println(s)

	i := 30
	s = i
	fmt.Println(s)

	e := Employee{
		Name: "Bob Bobson",
	}
	s = e
	fmt.Println(s)
}
