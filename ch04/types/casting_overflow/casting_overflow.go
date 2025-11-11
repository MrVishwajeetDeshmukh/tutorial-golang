package main

import (
	"fmt"
	"math"
)

func main() {
	// Assign the largest possible value for an int64 variable.
	// math.MaxInt64 = 9223372036854775807 (2^63 - 1)
	var int64Value int64 = math.MaxInt64

	// Convert (cast) int64 → int32 explicitly.
	// Since int32 can only hold values between -2,147,483,648 and 2,147,483,647,
	// this conversion causes an overflow.
	// The result will "wrap around" and give an incorrect (truncated) value.
	var int32Value int32 = int32(int64Value)

	// Print the results.
	// int64Value shows the correct large number.
	// int32Value shows the overflowed (wrapped) value.
	fmt.Printf("int32 value (after overflow): %d\n", int32Value)
	fmt.Printf("int64 value (original): %d\n", int64Value)
}
