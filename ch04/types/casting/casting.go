package main

import (
	"fmt"
	"math"
)

func main() {
	// Declare an int32 variable and assign it the maximum possible value for int32.
	// math.MaxInt32 = 2147483647 (2^31 - 1)
	var int32Value int32 = math.MaxInt32

	// Convert (cast) int32 to int64 explicitly.
	// This is necessary because Go does NOT perform implicit type conversions between integer types.
	// int64 can safely hold all values of int32 without overflow.
	var int64Value int64 = int64(int32Value)

	// Print both values. They will look identical, but the types are different.
	fmt.Printf("int32 value: %d\n", int32Value)
	fmt.Printf("int64 value: %d\n", int64Value)
}
