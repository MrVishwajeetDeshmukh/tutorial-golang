package main

import "fmt"

// Define Stringer interface
type Stringer interface {
	String() string
}

// Define a generic function
func PrintString[T Stringer](value T) {
	fmt.Println(value.String())
}

// Define a type that implements the Stringer interface
type Person struct {
	Name string
}

// Implement String method
func (p Person) String() string {
	return p.Name
}

func main() {
	// Call generic function using Person type
	p := Person{Name: "John Doe"}
	PrintString(p) // Output: John Doe
}
