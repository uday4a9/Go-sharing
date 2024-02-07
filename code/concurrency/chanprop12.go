package main

import "fmt"

func main() {
	ch1 := make(chan int)
	close(ch1) // HL
	ch1Ctr := 0

	ch2 := make(chan int)
	close(ch2) // HL
	ch2Ctr := 0

	for i := 0; i < 100; i++ {
		select { // HL
		case <-ch1:
			ch1Ctr++ // HL
		case <-ch2:
			ch2Ctr++ // HL
		}
	}

	fmt.Println("ch1Ctr:", ch1Ctr)
	fmt.Println("ch2Ctr:", ch2Ctr)
}
