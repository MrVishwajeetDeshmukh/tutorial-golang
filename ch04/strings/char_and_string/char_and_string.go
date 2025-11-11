package main

import (
	"fmt"
	"unsafe" // The 'unsafe' package allows checking the actual memory size of variables.
)

func main() {
	// 'byte' is an alias for 'uint8', meaning it can store a single ASCII character.
	var charA byte = 'a'

	// 'string' is a sequence of bytes (characters), immutable in Go.
	var stringA string = "a"
	var stringHello string = "hello world!"

	fmt.Println("format")
	fmt.Println("========")

	// When printing a byte with %v, it shows its numeric (ASCII) value.
	// 'a' in ASCII = 97
	fmt.Printf("byte charA: %v (number)\n", charA)

	// Using %c prints the byte as a character.
	fmt.Printf("byte charA: %c (char)\n", charA)

	// Using %s for a single byte prints an unexpected result,
	// because %s expects a string, not a single byte.
	fmt.Printf("byte charA: %s (string)\n\n", charA)

	// %s correctly prints a string.
	fmt.Printf("string stringA: %s (string)\n", stringA)

	// %c used with a string is invalid because a string is not a single character.
	// This prints the Unicode codepoint instead of a readable character.
	fmt.Printf("string stringA: %c (char)\n", stringA)

	// Strings can hold multiple characters.
	fmt.Printf("string stringHello: %v\n\n", stringHello)

	fmt.Println("size")
	fmt.Println("========")

	// 'byte' is a single 1-byte (8-bit) value.
	fmt.Printf("char size: %d bytes\n", unsafe.Sizeof(charA))

	// A 'string' in Go is actually a header structure (not the text itself).
	// It holds two fields:
	// 1. A pointer to the actual data
	// 2. The length of the string
	// That’s why it always takes 16 bytes on 64-bit systems.
	fmt.Printf("string size: %d bytes\n", unsafe.Sizeof(stringA))

	// Even large strings (longer text) occupy the same 16 bytes,
	// because 'string' only stores the pointer + length.
	fmt.Printf("large string size: %d bytes\n", unsafe.Sizeof(stringHello))
}
