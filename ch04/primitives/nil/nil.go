package main

import "fmt"

func main() {
	// Declare a slice of strings without initializing it
	var nilValue []string

	// Check if the slice is nil before assigning any value
	fmt.Printf("Is Nil?: %t\n", nilValue == nil)

	// Assign a value to the slice
	nilValue = []string{"hello"}

	// Check again if the slice is nil after assignment
	fmt.Printf("Is Nil?: %t\n", nilValue == nil)
}
