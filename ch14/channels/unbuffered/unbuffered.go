package main

import "fmt"

func main() {
	// Create an unbuffered channel
	unbufferedChannel := make(chan int)

	// Send value to the channel
	unbufferedChannel <- 1

	// Receive value from the channel
	fmt.Println(<-unbufferedChannel)
}
