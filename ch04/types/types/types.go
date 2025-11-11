package main

import "fmt"

func main() {
	// Define a constant for pi with double precision (float64).
	const pi float64 = 3.1415926535898932

	// Define a constant string for the name.
	const name string = "John Doe"

	// Age can never be negative, so we use an unsigned integer (uint8).
	var age uint8 = 28

	// Speed can be positive or negative (for direction), 
	// so we use a signed floating-point type (float32).
	var speed float32 = 60.5

	// Boolean variable — true means "is a citizen", false means "not a citizen".
	var isCitizen = true

	// Print pi using the %v format specifier (works for most types).
	fmt.Printf("The value of pi is %v\n", pi)
	fmt.Print("========\n")

	// Print a string (%s) and integer (%d).
	// %v could also be used for both if you don’t care about specific formatting.
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Speed is a floating-point value, so we use %f.
	// %f prints 6 decimal places by default.
	// %v prints more precise digits.
	fmt.Printf("Speed: %f\n", speed)

	// Simple conditional check using if-else.
	// If isCitizen is true → prints "You are a citizen."
	// Otherwise → prints "You are not a citizen."
	if isCitizen {
		fmt.Println("You are a citizen.")
	} else {
		fmt.Println("You are not a citizen.")
	}
}
