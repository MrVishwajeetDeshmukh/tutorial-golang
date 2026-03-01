package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu   sync.Mutex
	data int
)

// writer goroutine: periodically increments data value, terminates upon receiving stop signal
func writer(stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done() // Decrements WaitGroup counter on function exit
	for {
		select {
		case <-stop: // Received immediately if stop channel is closed (termination signal)
			fmt.Println("Writer: Stopping...")
			return // Terminate goroutine
		default: // Perform default action if there's no stop signal
			mu.Lock()
			data++
			fmt.Println("Writer: Incremented data to", data)
			mu.Unlock()
			time.Sleep(150 * time.Millisecond) // Simulate work interval
		}
	}
}

// reader goroutine: periodically reads data value, terminates upon receiving stop signal
func reader(stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done() // Decrements WaitGroup counter on function exit
	for {
		select {
		case <-stop: // Received immediately if stop channel is closed (termination signal)
			fmt.Println("Reader: Stopping...")
			return // Terminate goroutine
		default: // Perform default action if there's no stop signal
			mu.Lock()
			fmt.Println("Reader: Read data", data)
			mu.Unlock()
			time.Sleep(200 * time.Millisecond) // Simulate work interval
		}
	}
}

func main() {
	var wg sync.WaitGroup
	stop := make(chan struct{}) // Create channel for termination signal

	// Start writer and reader goroutines
	wg.Add(2) // Setup to wait for 2 goroutines
	go writer(stop, &wg)
	go reader(stop, &wg)

	// Wait for goroutines to run for a short duration
	time.Sleep(1 * time.Second) // Example: run for 1 second

	// Send termination signal to all goroutines
	fmt.Println("Main: Sending stop signal...")
	close(stop) // Sending signal to all listeners by closing the channel

	// Wait until all goroutines (writer, reader) have actually terminated
	fmt.Println("Main: Waiting for goroutines to stop...")
	wg.Wait()

	fmt.Println("Main: All tasks completed.")
}
