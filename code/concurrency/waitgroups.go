package main

import (
	"fmt"
	"sync"
	"time"
)

func ExecuteTask(taskID, timeToSleep int, wg *sync.WaitGroup) { // HL
	defer wg.Done() // HL
	fmt.Printf("Task %d started\n", taskID)
	time.Sleep(time.Duration(timeToSleep) * time.Second)
	fmt.Printf("Task %d completed\n", taskID)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3) // HL

	go ExecuteTask(1, 2, &wg)
	go ExecuteTask(2, 1, &wg)
	go ExecuteTask(3, 0, &wg)

	wg.Wait() // HL
	fmt.Println("All tasks completed")
}
