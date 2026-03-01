package main

import "fmt"

func main() {

	// Loop runs from 0 to 4 (total 5 iterations)
	for i := 0; i < 5; i++ {

		// If the number is even, skip the rest of this iteration
		if i%2 == 0 {
			continue
		}

		// This line runs only for odd numbers (1 and 3)
		fmt.Printf("%dth loop\n", i)
	}
}
