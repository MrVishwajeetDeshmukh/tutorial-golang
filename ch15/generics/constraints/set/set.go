package main

import "fmt"

// The Sum function defines a constraint for generic type T slices and sums
// T must be an integer or a floating-point number
func Sum[T int | float64](numbers []T) T {
	var total T
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	// Call the Sum function on slices of integers and floating-point numbers
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))     // Integer
	fmt.Println(Sum([]float64{1.1, 2.2, 3.3})) // Floating-point number
}
