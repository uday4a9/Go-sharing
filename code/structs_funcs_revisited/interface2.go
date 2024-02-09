package main

import "fmt"

type Employee struct {
	Name string
}

func receiver(s interface{}) int {
	// type switch
	switch v := s.(type) { // HL
	case string:
		return -1
	case Employee:
		return 0
	case int:
		return v
	default:
		fmt.Println("Unknown type")
		return -999
	}
}

func main() {
	i := 23
	s := "hello"
	e := Employee{
		Name: "Bob Bobson",
	}

	for _, v := range []interface{}{i, s, e, struct{}{}} {
		fmt.Println(receiver(v))
	}
}
