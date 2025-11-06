package main

import "fmt"

func main() {
	score := 100
	// → Declares a local variable named score and assigns it the value 100.
	// &score Returns the memory address of the variable score.

	fmt.Printf("Value stored in local variable 'score': %d\n", score)
	// %d Prints integers (so using %d to print an address may show an unexpected number or even cause a compiler warning).

	fmt.Printf("Memory address of local variable 'score' (as integer): %d\n", &score)
	//&score Returns the memory address of the variable score.

	fmt.Printf("Memory address of local variable 'score' (%%p): %p\n", &score)
	// %p Prints the pointer (address) in hexadecimal format — this is the correct way to display memory addresses.
}




