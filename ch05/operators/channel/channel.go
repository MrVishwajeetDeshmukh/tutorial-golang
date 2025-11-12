package main

import "fmt"

func main() {
	// Create a channel that can send and receive string values
	messages := make(chan string)

	// Start a goroutine (a lightweight thread)
	// This function sends the string "hello" into the channel
	go func() {
		messages <- "hello" // Send data into the channel
	}()

	// Receive the message from the channel
	// The program will wait (block) here until a value is sent
	msg := <-messages

	// Print the received message
	fmt.Println(msg) // Output: hello
}
