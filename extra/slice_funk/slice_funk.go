package main

import (
	"fmt"

	"github.com/thoas/go-funk"
)

func main() {
	original := []int{1, 2, 3, 4, 5}

	// Convert using funk.Map higher-order function
	doubled := funk.Map(original, func(value int) int {
		return value * 2
	}).([]int)

	fmt.Printf("Original slice: %v\n", original)
	fmt.Printf("Doubled slice: %v\n", doubled)
}
