package main

import "fmt"

func main() {
	// Declare two int32 variables
	var operand1 int32 = 29
	var operand2 int32 = 19

	// Perform subtraction and store the result
	var result int32 = operand1 - operand2

	// Print the result
	fmt.Printf("The result is: %d", result)
}

// Explanation:
// int32 specifies a 32-bit integer type.
// operand1 - operand2 performs a subtraction operation.
// %d is used to print integers in fmt.Printf.