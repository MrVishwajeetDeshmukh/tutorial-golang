package main

import "fmt"

// Outer function (function factory)
// It takes an integer 'x' and returns another function.
func funcFactory(x int) func(int) int {
	// Inner function
	// It takes another integer 'y' and returns the sum of x and y.
	return func(y int) int {
		return x + y
	}
}

func main() {
	// Create two different functions using funcFactory
	two := funcFactory(2)   // This function adds 2 to its input
	three := funcFactory(3) // This function adds 3 to its input

	// Call the returned functions
	fmt.Printf("Calling 'two' function gives: %d\n", two(10))
	fmt.Printf("Calling 'three' function gives: %d", three(10))
}
