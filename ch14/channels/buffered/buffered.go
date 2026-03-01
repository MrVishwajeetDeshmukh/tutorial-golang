package main

import "fmt"

func main() {
	// Create a buffered channel that can store up to 2 values
	channel := make(chan int, 2)

	// Send values to the channel
	channel <- 1
	channel <- 2

	// Close the channel
	close(channel)

	// Receive values from the channel
	fmt.Printf("First channel: %d\n", <-channel)
	fmt.Printf("Second channel: %d\n", <-channel)

	// Even if the channel is empty, retrieving a value will return the Zero value
	fmt.Printf("Third channel: %d\n", <-channel)
}
