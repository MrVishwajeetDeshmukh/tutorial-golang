package main

import "fmt"

// Function that takes two integers (a, b)
// and returns two integers: their sum and difference.
func addAndSub(a int, b int) (int, int) {
	add := a + b // calculate the sum
	sub := a - b // calculate the difference

	// return both values (sum, difference)
	return add, sub
}

func main() {
	// Define input values
	inputA := 2
	inputB := 3

	// Call addAndSub with inputA and inputB
	// Store the results in addResult and subResult
	addResult, subResult := addAndSub(inputA, inputB)

	// Print both results
	fmt.Printf("%v + %v = %v\n", inputA, inputB, addResult)
	fmt.Printf("%v - %v = %v\n", inputA, inputB, subResult)
}
