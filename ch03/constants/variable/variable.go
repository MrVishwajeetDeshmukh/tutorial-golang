package main

import "fmt"

func main() {
	// Changing 'const' to 'var' will still work without any issue.
	var pi = 3.1415926535898932
	var gravityConstant = 9.79641227572363

	// In fmt.Printf, this time variables (not constants) are inserted into the %v format specifier.
	fmt.Printf("The value of Pi is %v.\n", pi)
	fmt.Printf("The value of the gravitational constant G is %v.\n", gravityConstant)
}

// Explanation:
// var is used to declare variables, whose values can be changed later.
// The difference from const is that const values are fixed at compile time, while var values are mutable.
// %v is a format specifier that automatically chooses the correct format for the given type.