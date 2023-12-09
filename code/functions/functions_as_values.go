package main

import "fmt"

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

type opFuncType func(int, int) int // HL

var opMap = map[string]opFuncType{ // HL
	"+": add, "-": sub,
	"*": mul, "/": div,
}

func main() {
	a, b := 1, 2

	for k, v := range opMap {
		println(fmt.Sprintf("%d %s %d:", a, k, b), v(1, 2))
	}
}
