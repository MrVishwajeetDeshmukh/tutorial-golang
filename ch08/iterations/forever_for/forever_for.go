package main

import "fmt"

func main() {
	// Changed the condition so that it always evaluates to true.
	for i := 0; true; i++ {
		n := i + 1
		fmt.Printf("%dth iteration: i=%d\n", n, i)
	}
}
