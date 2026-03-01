package main

import "fmt"

// Define a generic function
func PrintLength[T any](value T) {
	// The code below only works if T is a type like string or []int.
	// However, in generics, such type checking cannot be guaranteed at compile time.
	fmt.Println(len(value)) // Original comment: This code requires T to support the len function.
}

func main() {
	// PrintLength("hello")        // Works normally

	// PrintLength([]int{1, 2, 3}) // Works normally
	// PrintLength(42) // Compilation error: int type does not support the len function.
}
