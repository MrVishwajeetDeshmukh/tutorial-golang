package main

import "fmt"

// Outer function that returns another function
func getIncreaser() func() {
	value := 100 // local variable inside the outer function

	// Inner function (a closure)
	return func() {
		// Access and modify 'value' from the outer function
		value += 50
		fmt.Printf("value: %d\n", value)
	}
}

func main() {
	// Call getIncreaser() which returns the inner function
	increaser := getIncreaser()

	// Each time we call 'increaser()', it keeps the previous value of 'value'
	for i := 0; i < 5; i++ {
		increaser() // prints increasing values each time
	}
}
