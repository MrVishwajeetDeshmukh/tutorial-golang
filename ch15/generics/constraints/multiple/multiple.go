package main

import "fmt"

// Define Stringer interface
type Stringer interface {
	String() string
}

// Define Error interface
type Error interface {
	Error() string
}

// Define generic type constraint
type StringerError interface {
	Stringer
	Error
}

// Define a generic function
func PrintInfo[T StringerError](value T) {
	fmt.Println(value.String())
	fmt.Println(value.Error())
}

// Define a type that implements Stringer and Error interfaces
type MyError struct {
	Message string
}

func (e MyError) String() string {
	return "String() => " + e.Message
}

func (e MyError) Error() string {
	return "Error() => " + e.Message
}

func main() {
	// Call generic function using MyError type
	e := MyError{Message: "Something went wrong"}
	PrintInfo(e)
}
