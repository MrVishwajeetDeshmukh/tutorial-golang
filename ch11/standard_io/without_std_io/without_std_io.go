package main

import (
	"fmt"
	"os"
)

func main() {
	// Output directly to the current terminal
	// Open the tty session file directly for output
	file, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Error opening /dev/tty:", err)
		return
	}
	defer file.Close()

	// Write a message to the terminal
	_, err = file.WriteString("Hello, this is a message sent directly to the terminal!\n")
	if err != nil {
		fmt.Println("Error writing to terminal:", err)
		return
	}
}
