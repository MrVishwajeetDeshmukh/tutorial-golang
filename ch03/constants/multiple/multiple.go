package main

import "fmt"

func main() {
	// You can declare multiple constants using parentheses.
	const (
		name = "Hong Gil-dong"
		age  = 31
	)

	// Multiple constants can also be declared on a single line.
	const pi, gravityConstant = 3.1415926535898932, 9.79641227572363

	fmt.Printf("Hello, my name is %v and I am %d years old.\n", name, age)
	fmt.Printf("The value of Pi is %v.\n", pi)
	fmt.Printf("The value of the gravitational constant G is %v.\n", gravityConstant)
}

// Explanation:
// You can group constants inside parentheses for readability.
// You can also declare multiple constants on the same line.
// %v → prints the value in a default format.
// %d → prints an integer value (used here for age).

