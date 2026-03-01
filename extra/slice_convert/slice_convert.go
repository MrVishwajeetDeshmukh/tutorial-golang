package main

import "fmt"

func main() {
	original := []int{1, 2, 3, 4, 5}

	// Allocate a slice of the same size as the original
	// to store the converted values
	doubled := make([]int, len(original))

	// Write the conversion code
	for i, value := range original {
		doubled[i] = value * 2
	}

	fmt.Printf("Original slice: %v\n", original)
	fmt.Printf("Doubled slice: %v\n", doubled)
}
