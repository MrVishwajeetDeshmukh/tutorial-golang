package main

import "fmt"

func checkConditionWithBool() {
	// This condition is always true, so the code inside this block will run.
	if true {
		fmt.Println("Checked true!")
	} else {
		// This block will never run because the condition above is always true.
		fmt.Println("Check false!")
	}
}

func main() {
	checkConditionWithBool()
}
