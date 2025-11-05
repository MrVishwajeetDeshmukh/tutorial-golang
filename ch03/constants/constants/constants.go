package main

import "fmt"

func main() {
	// 'const' is a keyword used to declare constants.
	const pi = 3.1415926535898932
	const gravityConstant = 9.79641227572363

	// In fmt.Printf, the constants are inserted into the %v format specifier for output.
	fmt.Printf("The value of Pi is %v.\n", pi)
	fmt.Printf("The value of the gravitational constant G is %v.\n", gravityConstant)
}

