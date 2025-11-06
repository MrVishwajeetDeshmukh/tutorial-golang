package main

import "fmt"

func main() {
	// Declare a pointer variable of type int.
	// A pointer stores the memory address of another variable.
	var ptr *int

	// Declare and initialize an int variable.
	var val int = 42

	// Assign the address of 'val' to 'ptr'.
	// Now 'ptr' points to the memory location where 'val' is stored.
	ptr = &val

	// Dereference 'ptr' to access the value it points to.
	// '*ptr' gives the actual value stored at that memory address.
	fmt.Printf("Value pointed to by ptr: %d\n", *ptr)

	// Change the value stored at the memory address 'ptr' points to.
	// This will also update the original 'val' variable.
	*ptr = 100

	// Print 'val' to verify that its value was changed through the pointer.
	fmt.Printf("New value of val: %d\n", val)
}
