package main

func main() {
	// created a channel and closed it
	ch := make(chan int)

	go func() {
		defer close(ch) // HL
		ch <- 12
	}()

	val, ok := <-ch             // HL
	println("READ 1:", val, ok) // HL

	val, ok = <-ch              // HL
	println("READ 2:", val, ok) // HL
}
