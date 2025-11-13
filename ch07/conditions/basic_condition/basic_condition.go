package main

import "fmt"

// getScoreRank returns a grade (A, B, C, F)
// based on the score provided.
func getScoreRank(score int) (rank string) {
	if score >= 90 {
		return "A" // 90 or above → A grade
	} else if score >= 80 {
		return "B" // 80–89 → B grade
	} else if score >= 70 {
		return "C" // 70–79 → C grade
	} else {
		return "F" // below 70 → F grade
	}
}

func main() {
	var score int

	// Ask the user to enter their score
	fmt.Print("Enter your score: ")
	fmt.Scanf("%d\n", &score)

	// Call the function to get the rank based on the score
	rank := getScoreRank(score)

	// Print the grade result
	fmt.Printf("Your grade is '%v'.\n", rank)
}
