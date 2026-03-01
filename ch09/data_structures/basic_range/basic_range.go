package main

import "fmt"

func main() {
	score := make([]int, 5)

	// Ask the user to input scores for 5 subjects
	for i := range score {
		fmt.Printf("Enter the score for subject %d: ", i+1)
		fmt.Scanf("%d\n", &score[i])
	}
	fmt.Println("============================")

	// Iterate through all slice elements and calculate the average score
	sumScore := 0
	for _, scoreItem := range score {
		sumScore += scoreItem
	}
	meanScore := float64(sumScore) / float64(len(score))
	fmt.Printf("Your average score is %.2f.\n", meanScore)
}
