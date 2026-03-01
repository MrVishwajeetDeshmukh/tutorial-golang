package main

import "fmt"

// Generic function with constraints that satisfying both Stringer and Error interfaces
func PrintDetails[T interface {
	fmt.Stringer
	error
}](value T) {
	fmt.Println("String:", value.String())
	fmt.Println("Error:", value.Error())
}

// Type that implements Stringer and Error interfaces
type MyError struct {
	Message string
}

func (e MyError) String() string {
	return e.Message
}

func (e MyError) Error() string {
	return e.Message
}

func main() {
	// Call generic function using MyError type
	e := MyError{Message: "An error occurred"}
	PrintDetails(e)
}
