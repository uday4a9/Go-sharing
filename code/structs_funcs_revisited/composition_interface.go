package main

import (
	"fmt"
)

// Animal struct representing basic properties of an animal
type Animal struct {
	Name  string
	Color string
}

// Dog struct representing a specific type of animal
type Dog struct {
	Animal // Embedding Animal struct for composition
	Breed  string
}

// Describe method for Dog struct to describe the dog
func (d Dog) Describe() { // HL
	fmt.Printf("This is a %s dog named %s\n", d.Breed, d.Name)
}

// Cat struct representing another specific type of animal
type Cat struct {
	Animal // Embedding Animal struct for composition
	Sound  string
}

// Describe method for Cat struct to describe the cat
func (c Cat) Describe() {
	fmt.Printf("This is a cat named %s that says %s\n", c.Name, c.Sound)
}

type Describer interface { // HL
	// Describe method for Describer interface
	Describe() // HL
}

func describeAnimal(d Describer) { // HL
	d.Describe() // HL
}

func main() {
	var myDog Describer // HL

	// Creating a Dog instance and assign it to Describer interface instance
	myDog = Dog{
		Animal: Animal{Name: "Buddy", Color: "Brown"},
		Breed:  "Labrador",
	}

	// Describe the dog and cat
	describeAnimal(myDog)
}
