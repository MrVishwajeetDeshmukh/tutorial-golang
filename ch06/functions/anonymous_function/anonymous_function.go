package main

import "fmt"

// A normal (named) function with a named return value 'sum'
// It takes two integers (a and b) and returns their sum.
func namedFunction(a int, b int) (sum int) {
	return a + b
}

func main() {
	// Call the named function and store the result
	namedAddResult := namedFunction(2, 3)

	// Define and call an anonymous function (a function without a name)
	// directly in one line using (2, 3) as arguments.
	anonymousAddResult := func(a int, b int) (sum int) {
		return a + b
	}(2, 3)

	// Print both results
	fmt.Printf("namedAdd result: %v\n", namedAddResult)
	fmt.Printf("anonymousAdd result: %v\n", anonymousAddResult)
}
