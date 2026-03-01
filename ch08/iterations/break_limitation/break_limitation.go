package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ { // Outer loop runs 5 times (i = 0 to 4)
		for j := 0; j < 3; j++ { // Inner loop runs 3 times (j = 0 to 2)
			fmt.Printf("i=%d, j=%d\n", i, j)

			// When j becomes 1, stop only the inner loop.
			// The outer loop continues with the next i.
			if j == 1 {
				break
			}
		}
	}
}
