package main

import "fmt"

func main() {
	fmt.Println("Move to line 10")
	goto inside_loop

	for i := 0; i < 5; i++ {
	inside_loop:
		fmt.Printf("%dth iteration\n", i)
	}
}
