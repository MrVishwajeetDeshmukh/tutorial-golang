package main

import "fmt"

func main() {
	// Declare and initialize a variable
	var x int = 10

	// Reference operator (&)
	// The '&' operator gives the memory address of a variable
	var p *int = &x
	fmt.Println(p) // Prints the memory address of x

	// Dereference operator (*)
	// The '*' operator gives the value stored at the memory address
	fmt.Println(*p) // Prints the value of x (10)

	// Modify the value of x using the pointer
	*p = 20
	fmt.Println(x) // Prints 20 (x is updated through the pointer)
}
