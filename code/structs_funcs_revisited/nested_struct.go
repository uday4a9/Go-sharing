package main

import "fmt"

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s, %s, %s", a.Street, a.City, a.State, a.Zip)
}

type Employee struct {
	Name    string
	ID      string
	Address Address // HL
}

func (e Employee) Description() {
	fmt.Printf("%s (%s)\n", e.Name, e.ID)
	fmt.Println(e.Address.FullAddress()) // HL
}

func main() {
	e := Employee{
		Name: "Bob Bobson",
		ID:   "12345",
		Address: Address{
			Street: "123 Main St",
			City:   "Springfield",
			State:  "IL",
			Zip:    "62701",
		},
	}
	e.Description() // HL
}
