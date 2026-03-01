package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		break
		fmt.Printf("%dth loop\n", i)
	}
}
