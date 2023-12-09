package main

func gen(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i * 2
	}
	// close(ch)
}

func main() {
	ch := make(chan int)
	go gen(ch)
	for i := range ch {
		println(i)
	}
}
