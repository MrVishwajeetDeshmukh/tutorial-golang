package main

import (
	"fmt"
)

func main() {
	var name string

	// Get name from the user
	fmt.Print("Please enter your name: ")
	fmt.Scanf("%s\n", &name)

	// Print the entered name
	fmt.Printf("Hello, %s!\n", name)
}
