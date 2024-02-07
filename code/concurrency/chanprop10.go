package main

func main() {
	// created a channel and closed it
	ch := make(chan int)
	close(ch) // HL

	val, ok := <-ch             // HL
	println("READ 1:", val, ok) // HL

	val, ok = <-ch              // HL
	println("READ 2:", val, ok) // HL
}
