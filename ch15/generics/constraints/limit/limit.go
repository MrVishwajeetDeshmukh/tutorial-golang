package main

import "fmt"

// Define a generic function without constraints
func PrintLength[T any](value T) {
	fmt.Println(len(value)) // This code requires T to support the len function.
}

func main() {
	PrintLength("hello")        // Works normally
	PrintLength([]int{1, 2, 3}) // Works normally
	// PrintLength(42) // Compilation error: int type does not support the len function.
}
