package main

import "fmt"

// Define a function named 'add' that takes two integers (a and b)
// and returns their sum as an integer.
func add(a int, b int) int {
	return a + b
}

func main() {
	// Call the add() function with 2 and 3 as arguments
	// and store the result in the 'result' variable.
	result := add(2, 3)

	// Print the result to the console.
	fmt.Printf("result: %v\n", result) // Output: result: 5
}
