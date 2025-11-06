package main

import "fmt"

// updateValue receives an integer value (not a pointer).
// Since Go uses "pass by value", the function only gets a copy of the original value.
func updateValue(value int) {
	value = 100
}

func main() {
	myValue := 50
	fmt.Printf("Before change: %d\n", myValue)

	// Passes the value of myValue, not its address.
	updateValue(myValue)

	fmt.Printf("After change: %d\n", myValue)
}
