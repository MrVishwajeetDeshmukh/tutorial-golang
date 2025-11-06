package main

import "fmt"

// 'value' is a pointer parameter that receives the memory address of an int variable.
// This kind of variable is called a *pointer*.
func updateValue(value *int) {
	// Dereference the pointer and update the value stored at that memory address to 100.
	*value = 100
}

func main() {
	myValue := 50
	fmt.Printf("Before change: %d\n", myValue)

	// Pass the memory address of 'myValue' to the function.
	updateValue(&myValue)

	fmt.Printf("After change: %d\n", myValue)
}
