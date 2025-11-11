package main

import "fmt"

func main() {
	// 'byte' is an alias for uint8, which can store values from 0 to 255 (1 byte = 8 bits).
	// The character below uses a Unicode value that exceeds 255, which cannot fit into a byte.
	// This will cause a compile-time error: "constant ... overflows byte".

	var byteChar byte = 'A' // ✅ valid example (fits in one byte)
	// var byteChar byte = '😀' // ❌ would cause overflow (Unicode value > 255)

	// Print the numeric (ASCII) value of the byte.
	fmt.Printf("byte char (numeric value): %v\n", byteChar)

	// Print the character representation of the byte using %c.
	fmt.Printf("byte char (as character): %c\n", byteChar)
}
