package main

import (
	"fmt"
	"time"
)

// Goroutine that performs a long-running task
func longRunningTask() {
	for {
		// Endless loop execution
		fmt.Println("Long-running task is still running...")
		time.Sleep(1 * time.Second)
	}
}

func startTask() {
	go longRunningTask() // Start goroutine performing long-running task
}

func main() {
	startTask() // Start goroutine performing long-running task in another function

	// Main function will end after 2 seconds
	time.Sleep(2 * time.Second)
	fmt.Println("Main function ends.")
}
