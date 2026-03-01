package main

import "fmt"

func main() {
	// Literal slice with length 6
	original := []int{1, 2, 3, 4, 5, 6}

	// Slicing values from the second to the fourth element
	slicing := original[1:4]

	fmt.Println("================================================")
	fmt.Println("[original]")
	fmt.Printf("len: %d, cap: %d\n", len(original), cap(original))
	fmt.Printf("value: %v\n", original)
	fmt.Println("================================================")
	fmt.Print("\n")

	fmt.Println("================================================")
	fmt.Println("[slicing]")
	fmt.Printf("len: %d, cap: %d\n", len(slicing), cap(slicing))
	fmt.Printf("value: %v\n", slicing)
	fmt.Println("================================================")
}
