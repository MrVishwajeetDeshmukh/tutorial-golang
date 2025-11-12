package main

import "fmt"

func main() {
	// Two unsigned 8-bit integers (uint8)
	var value1 uint8 = 187 // binary: 10111011
	var value2 uint8 = 222 // binary: 11011110

	// Value to use for shift operations
	var shiftValue uint8 = 19 // binary: 00010011

	// Number of bit positions to shift
	var shiftLength uint8 = 2

	// -----------------------------
	// AND Bitwise Operation
	// -----------------------------
	// Only bits that are 1 in both values remain 1
	fmt.Println("AND result")
	fmt.Println("========")
	andResult := value1 & value2 // 10011010 (decimal 154)
	fmt.Printf("%08b & %08b = %08b\n", value1, value2, andResult)

	// -----------------------------
	// OR Bitwise Operation
	// -----------------------------
	// Bits that are 1 in either value become 1
	fmt.Print("\n")
	fmt.Println("OR result")
	fmt.Println("========")
	orResult := value1 | value2 // 11111111 (decimal 255)
	fmt.Printf("%08b | %08b = %08b\n", value1, value2, orResult)

	// -----------------------------
	// NOT Bitwise Operation
	// -----------------------------
	// Flips all bits (1 → 0, 0 → 1)
	fmt.Print("\n")
	fmt.Println("NOT result")
	fmt.Println("========")
	notResult1 := ^value1 // 01000100 (decimal 68)
	notResult2 := ^value2 // 00100001 (decimal 33)
	fmt.Printf("^%08b: %08b\n", value1, notResult1)
	fmt.Printf("^%08b: %08b\n", value2, notResult2)

	// -----------------------------
	// XOR Bitwise Operation
	// -----------------------------
	// Bits that are different become 1
	fmt.Print("\n")
	fmt.Println("XOR result")
	fmt.Println("========")
	xorResult := value1 ^ value2 // 01100101 (decimal 101)
	fmt.Printf("%08b ^ %08b = %08b\n", value1, value2, xorResult)

	// -----------------------------
	// AND NOT Bitwise Operation
	// -----------------------------
	// Keeps bits from the first value that are 0 in the second
	fmt.Print("\n")
	fmt.Println("AND NOT result")
	fmt.Println("========")
	andNotResult := value1 &^ value2 // 00100001 (decimal 33)
	fmt.Printf("%08b &^ %08b = %08b\n", value1, value2, andNotResult)

	// -----------------------------
	// Left Shift Operation
	// -----------------------------
	// Moves bits to the left (× 2 for each shift)
	fmt.Print("\n")
	fmt.Println("Left shift result")
	fmt.Println("========")
	leftShiftResult := shiftValue << shiftLength // 01001100 (decimal 76)
	fmt.Printf("%08b << %08b = %08b\n", shiftValue, shiftLength, leftShiftResult)

	// -----------------------------
	// Right Shift Operation
	// -----------------------------
	// Moves bits to the right (÷ 2 for each shift)
	fmt.Print("\n")
	fmt.Println("Right shift result")
	fmt.Println("========")
	rightShiftResult := shiftValue >> shiftLength // 00000100 (decimal 4)
	fmt.Printf("%08b >> %08b = %08b\n", shiftValue, shiftLength, rightShiftResult)
}
