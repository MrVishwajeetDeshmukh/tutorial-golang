package main

import "fmt"

// Define a generic type
type Pair[T any] struct {
	First  T
	Second T
}

func main() {
	// Integer type Pair
	intPair := Pair[int]{1, 2}
	fmt.Println("Integer Pair:", intPair) // Output: Integer Pair: {1 2}

	// String type Pair
	strPair := Pair[string]{"hello", "world"}
	fmt.Println("String Pair:", strPair) // Output: String Pair: {hello world}
}
