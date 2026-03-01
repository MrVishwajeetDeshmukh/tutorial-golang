package main

import "fmt"

// Send-only unidirectional channel
func sendChannel(value int, channel chan<- int) {
	channel <- value
}

// Receive-only unidirectional channel
func receiveChannel(channel <-chan int) {
	fmt.Printf("Channel receive: %d\n", <-channel)
}

func main() {
	// Create a bidirectional channel
	channel := make(chan int, 1)

	// Send value
	sendChannel(1, channel)

	// Receive value
	receiveChannel(channel)
}
