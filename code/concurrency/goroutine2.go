package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("Hello")
}

func main() {
	go printHello() // HL
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
