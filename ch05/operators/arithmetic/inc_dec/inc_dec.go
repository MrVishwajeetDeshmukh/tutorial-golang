package main

import "fmt"

func main() {
	// Start with 1
	increase := 1

	// The ++ operator adds 1 to the current value
	// Same as: increase = increase + 1
	increase++
	fmt.Println(increase) // prints 2

	// Start with 1 again
	decrease := 1

	// The -- operator subtracts 1 from the current value
	// Same as: decrease = decrease - 1
	decrease--
	fmt.Println(decrease) // prints 0
}
