package main

import "fmt"

// Define a generic function
func Swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {
	// Swap integer values
	x, y := 1, 2
	x, y = Swap(x, y)
	fmt.Println("Swapped:", x, y) // Output: Swapped: 2 1

	// Swap string values
	str1, str2 := "hello", "world"
	str1, str2 = Swap(str1, str2)
	fmt.Println("Swapped:", str1, str2) // Output: Swapped: world hello
}
