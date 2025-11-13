package main

import "fmt"

// getScoreRank returns a grade (A, B, C, F)
// based on the score using a switch statement.
func getScoreRank(score int) (rank string) {
	switch {
	case score >= 90:
		return "A" // 90 or above → A grade
	case score >= 80:
		return "B" // 80–89 → B grade
	case score >= 70:
		return "C" // 70–79 → C grade
	default:
		return "F" // below 70 → F grade
	}
}

func main() {
	var score int

	// Ask the user to enter their score
	fmt.Print("Enter your score: ")
	fmt.Scanf("%d\n", &score)

	// Call the function to get the grade
	rank := getScoreRank(score)

	// Print the result
	fmt.Printf("Your grade is '%v'.\n", rank)
}
