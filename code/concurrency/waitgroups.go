package main

import (
	"fmt"
	"sync"
	"time"
)

func doWait(duration time.Duration) {
	time.Sleep(duration)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		doWait(2 * time.Second)
		fmt.Println("Task 1 completed")
	}()
	go func() {
		defer wg.Done()
		doWait(1 * time.Second)
		fmt.Println("Task 2 completed")
	}()
	go func() {
		defer wg.Done()
		doWait(0 * time.Second)
		fmt.Println("Task 3 completed")
	}()
	wg.Wait()
	fmt.Println("All tasks completed")
}
