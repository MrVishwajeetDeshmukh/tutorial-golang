package main

import "fmt"

func main() {
	// Declare a nil slice for the variable sliceValue
	var sliceValue []int

	// Add 4 elements to the sliceValue variable
	sliceValue = append(sliceValue, 1, 2, 3, 4)

	// Print the value of the sliceValue variable
	fmt.Printf("sliceValue len: %d, cap: %d, value: %v\n",
		len(sliceValue), cap(sliceValue), sliceValue)
}
