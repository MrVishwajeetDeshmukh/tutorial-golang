package main

import "fmt"

func sendValue(value int, channel chan int) {
	channel <- value
}

func main() {
	// Create an unbuffered channel
	unbufferedChannel := make(chan int)

	// Send value to the channel using a goroutine
	go sendValue(1, unbufferedChannel)

	// Receive value from the channel
	fmt.Println(<-unbufferedChannel)
}
