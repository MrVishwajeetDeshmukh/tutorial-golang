package main

import (
	"fmt"
	"strings"
)

func main() {
	original := "hello world"
	separator := " "
	result := strings.SplitN(original, separator, 2)

	fmt.Printf("Original string: '%s'\n", original)
	fmt.Printf("Separator string: '%s'\n", separator)

	fmt.Println("================================================")
	fmt.Printf("Separated string slice: %v\n", result)
}
