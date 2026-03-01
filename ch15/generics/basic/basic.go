package main

import "fmt"

// Define a generic type
type Pair[T any] struct {
	First  T
	Second T
}

// Define a generic function
func PrintPair[T any](p Pair[T]) {
	fmt.Printf("First: %v, Second: %v\n", p.First, p.Second)
}

func main() {
	// Pair using integer type
	intPair := Pair[int]{1, 2}
	PrintPair(intPair)

	// Pair using string type
	strPair := Pair[string]{"hello", "world"}
	PrintPair(strPair)
}
