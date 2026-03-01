package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a sentence: ")
	text, _ := reader.ReadString('\n') // Read input including newline character
	text = strings.TrimSpace(text)     // Remove unnecessary whitespace
	fmt.Printf("Entered sentence: %s\n", text)
}
