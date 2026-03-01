package main

import (
	"fmt"
	"sync"
	"time"
)

// Goroutine that performs a long-running task
func longRunningTask(stop chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			// Terminate loop upon receiving termination signal
			fmt.Println("Long-running task is stopping...")
			return
		default:
			// Endless loop execution
			fmt.Println("Long-running task is still running...")
			time.Sleep(1 * time.Second)
		}
	}
}

func startTask(stop chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go longRunningTask(stop, wg) // Start goroutine performing long-running task
}

func main() {
	var wg sync.WaitGroup
	stop := make(chan struct{})

	startTask(stop, &wg) // Start goroutine performing long-running task in another function

	// Main function will end after 2 seconds
	time.Sleep(2 * time.Second)
	fmt.Println("Main function ends.")

	// Send termination signal to all goroutines
	close(stop)

	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("All tasks completed")
}
