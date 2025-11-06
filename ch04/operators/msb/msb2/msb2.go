package main

import "fmt"

func main() {
	// Declare operand1 as 29 (a positive number)
	var operand1 int32 = 29

	// Use bitwise NOT (~) and add 1 to get the two's complement of 19
	// This effectively represents -19
	var operand2 int32 = ^19 + 1

	// Perform addition: 29 + (-19)
	var result int32 = operand1 + operand2

	// Print operand1 in 32-bit binary format
	fmt.Printf("operand1: %.32b\n", operand1)

	// Print operand2 in 32-bit binary format
	fmt.Printf("operand2: %.32b\n", operand2)

	// Print result in 32-bit binary format
	fmt.Printf("Result (binary): %.32b\n", result)

	// Print result as a decimal value
	fmt.Printf("Result (decimal): %d", result)
}
